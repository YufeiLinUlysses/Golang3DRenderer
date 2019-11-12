package feature

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

/*OBJParser type contains all necessary compnent of a objparser
 *OBJParser contains a slice of tuples*/
type OBJParser struct {
	FilePath string
	Vertices []Tuple
	Groups   map[string]*Group
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
		converted, ispoint, isgroup, points, gname, finErr := convertLine(scanner.Text())
		if finErr != nil {
			WriteErrorFile(errorStr, finErr)
		}
		if !converted {
			continue
		}
		if ispoint {
			op.Vertices = append(op.Vertices, *Point(points[0], points[1], points[2]))
		}
		if isgroup {
			tempGname = gname
			if _, ok := op.Groups[gname]; ok {
				continue
			} else {
				op.Groups[gname] = NewGroup()
			}
		}
		if !ispoint && !isgroup {
			tris := fanTriangulation(op.Vertices, points)
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
 *convertLine returns three bool, a slice of float64, a string and an error*/
func convertLine(line string) (converted, ispoint, isgroup bool, point []float64, groupname string, finErr error) {
	var cors []float64
	splitted := strings.Split(line, " ")
	if splitted[0] == "v" && len(splitted) != 4 {
		return false, false, false, point, "", nil
	} else if splitted[0] == "v" {
		for i := 1; i <= 3; i++ {
			if float, err := strconv.ParseFloat(splitted[i], 64); err == nil {
				cors = append(cors, float)
			} else {
				return false, false, false, point, "", err
			}
		}
		return true, true, false, cors, "", nil
	}
	if splitted[0] == "f" {
		for i := 1; i < len(splitted); i++ {
			if float, err := strconv.ParseFloat(splitted[i], 64); err == nil {
				cors = append(cors, float)
			} else {
				return false, false, false, point, "", err
			}
		}
		return true, false, false, cors, "", nil
	}
	if splitted[0] == "g" {
		return true, false, true, cors, line[2:], nil
	}
	return false, false, false, point, "", finErr
}

/*fanTriangulation interprets the obj file and recreate a complicate shape with triangles
 *fanTriangulation takes in a slice of Tuple, and a slice of float64
 *fanTriangulation returns a slice of interface{}*/
func fanTriangulation(vertices []Tuple, point []float64) []interface{} {
	var tris []interface{}
	for i := 1; i < len(point)-1; i++ {
		tris = append(tris, *NewTriangle(&vertices[int(point[0])-1], &vertices[int(point[i])-1], &vertices[int(point[i+1])-1]))
	}
	return tris
}
