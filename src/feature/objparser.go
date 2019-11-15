package feature

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*OBJParser type contains all necessary compnent of a objparser
 *OBJParser contains a slice of tuples*/
type OBJParser struct {
	FilePath string
	Vertices []Tuple
	Normals  []Tuple
	Groups   map[string]*Group
}

/*decide type contains all necessary component of a decide*/
type decide struct {
	converted                       bool
	ispoint, isgroup, isnormal, isf bool
	points                          []float64
	position                        []int
	normals                         []float64
	normalpos                       []int
	groupname                       string
	finErr                          error
}

/*NewOBJParser creates an instance of Type OBJParser
 *NewOBJParser takes in a string
 *NewOBJParser returns a sphere with default object*/
func NewOBJParser(title string) *OBJParser {
	grs := make(map[string]*Group)
	op := &OBJParser{
		FilePath: title,
		Groups:   grs,
	}
	return op
}

/*newdecide creates an instance of Type decide
 *newdecide returns a decide*/
func newdecide() *decide {
	po := make([]float64, 3)
	pos := make([]int, 0)
	d := &decide{
		converted: false,
		ispoint:   false,
		isgroup:   false,
		isnormal:  false,
		isf:       false,
		points:    po,
		position:  pos,
		normals:   po,
		normalpos: pos,
		groupname: "",
	}
	return d
}

/*ReadObj reads the obj file and converts it to an OBJParser instance
 *ReadObj can only be called by a OBJParser instance
 *ReadObj returns an OBJParser*/
func (op *OBJParser) ReadObj() *OBJParser {
	errorStr := "Writing File:" + op.FilePath + time.Now().String() + "\n"
	tempGname := ""
	file, err := os.Open(op.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		d := convertLine(scanner.Text())
		if d.finErr != nil {
			WriteErrorFile(errorStr, d.finErr)
		}
		if !d.converted {
			continue
		}
		if d.ispoint {
			op.Vertices = append(op.Vertices, *Point(d.points[0], d.points[1], d.points[2]))
		}
		if d.isgroup {
			tempGname = d.groupname
			if _, ok := op.Groups[d.groupname]; ok {
				continue
			} else {
				op.Groups[d.groupname] = NewGroup()
			}
		}
		if d.isnormal {
			op.Normals = append(op.Normals, *Vector(d.normals[0], d.normals[1], d.normals[2]))
		}
		if d.isf {
			hasNormal := len(op.Normals) > 0
			tris := fanTriangulation(op.Vertices, op.Normals, d.position, d.normalpos, hasNormal)
			for i := range tris {
				op.Groups[tempGname].AddChild(tris[i])
			}
		}
	}
	if err := scanner.Err(); err != nil {
		WriteErrorFile(errorStr, err)
		return nil
	}
	return op
}

/*OBJToGroup converts OBJParser to Group
 *OBJToGroup can only be called by an OBJParser
 *OBJToGroup returns a Group*/
func (op *OBJParser) OBJToGroup() *Group {
	var gr Group
	for key := range op.Groups {
		gr.AddChild(op.Groups[key])
	}
	return &gr
}

/*convertLine converts a string to things we need for justification in converting obj file
 *convertLine takes in a string
 *convertLine returns a decide*/
func convertLine(line string) *decide {
	d := newdecide()
	//regex
	rev := regexp.MustCompile(`v ([-+]?[0-9]+[\.]?[0-9]* ?){3}`)
	reg := regexp.MustCompile(`g \w+`)
	ref1 := regexp.MustCompile(`f (\d+ ?){3,}`)
	ref2 := regexp.MustCompile(`f (\d+//\d+ ?){3,}`)
	ref3 := regexp.MustCompile(`f (\d+/\d+/\d+ ?){3,}`)
	revn := regexp.MustCompile(`vn ([-+]?[0-9]+[\.]?[0-9]* ?){3}`)

	splitted := strings.Split(line, " ")

	switch {
	case rev.MatchString(line):
		if len(splitted) == 4 {
			d.converted = true
			d.ispoint = true
			for i := 1; i <= 3; i++ {
				float, _ := strconv.ParseFloat(splitted[i], 64)
				d.points[i-1] = float
			}
		}
	case reg.MatchString(line):
		d.converted = true
		d.isgroup = true
		d.groupname = line[2:]
	case revn.MatchString(line):
		if len(splitted) == 4 {
			d.converted = true
			d.isnormal = true
			for i := 1; i <= 3; i++ {
				float, _ := strconv.ParseFloat(splitted[i], 64)
				d.normals[i-1] = float
			}
		}
	case ref1.MatchString(line):
		d.converted = true
		d.isf = true
		for i := 1; i < len(splitted); i++ {
			float, _ := strconv.ParseFloat(splitted[i], 64)
			d.position = append(d.position, int(float))
		}
	case ref2.MatchString(line):
		d.converted = true
		d.isf = true
		for i := 1; i < len(splitted); i++ {
			temp := strings.Split(splitted[i], "//")
			float1, _ := strconv.ParseFloat(temp[0], 64)
			d.position = append(d.position, int(float1))
			float2, _ := strconv.ParseFloat(temp[1], 64)
			d.normalpos = append(d.normalpos, int(float2))
		}
	case ref3.MatchString(line):
		d.converted = true
		d.isf = true
		for i := 1; i < len(splitted); i++ {
			temp := strings.Split(splitted[i], "/")
			float1, _ := strconv.ParseFloat(temp[0], 64)
			d.position = append(d.position, int(float1))
			float2, _ := strconv.ParseFloat(temp[2], 64)
			d.normalpos = append(d.normalpos, int(float2))
		}
	}
	return d
}

/*fanTriangulation interprets the obj file and recreate a complicate shape with triangles or smoothtriangles
 *fanTriangulation takes in two slices of Tuple, two slices of int and a bool
 *fanTriangulation returns a slice of interface{}*/
func fanTriangulation(vertices, normals []Tuple, position, normalpos []int, hasNormal bool) []interface{} {
	var tris []interface{}
	if hasNormal {
		for i := 1; i < len(position)-1; i++ {
			tris = append(tris, *NewSmoothTriangle(vertices[0], vertices[position[i]-1], vertices[position[i+1]-1],
				normals[normalpos[0]-1], normals[normalpos[i]-1], normals[normalpos[i+1]-1]))
		}
	} else {
		for i := 1; i < len(position)-1; i++ {
			tris = append(tris, *NewTriangle(&vertices[0], &vertices[position[i]-1], &vertices[position[i]-1]))
		}
	}
	return tris
}
