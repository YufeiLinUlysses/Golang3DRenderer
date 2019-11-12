package feature

import (
	"bufio"
	"fmt"
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
	DefGroup *Group
	Vertices []Tuple
}

/*NewOBJParser initiates*/
func NewOBJParser(title string) *OBJParser {
	op := &OBJParser{
		FilePath: title,
		DefGroup: NewGroup(),
	}
	return op
}

/*ReadObj reads*/
func (op *OBJParser) ReadObj() *OBJParser {
	errorStr := "Writing File:" + op.FilePath + time.Now().String() + "\n"
	file, err := os.Open(op.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		converted, ispoint, points, finErr := convertLine(scanner.Text())
		if finErr != nil {
			WriteErrorFile(errorStr, finErr)
		}
		if !converted {
			continue
		}
		if ispoint {
			op.Vertices = append(op.Vertices, *Point(points[0], points[1], points[2]))
		} else {
			tris := fanTriangulation(op.Vertices, points)
			for i := range tris {
				op.DefGroup.AddChild(tris[i])
			}
		}
	}
	fmt.Println(len(op.DefGroup.Objects))
	fmt.Println(op.DefGroup.Objects[0].(Triangle).Parent)
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		WriteErrorFile(errorStr, err)
		return nil
	}
	return op
}

/*convertLine converts*/
func convertLine(line string) (converted, ispoint bool, point []float64, finErr error) {
	var cors []float64
	splitted := strings.Split(line, " ")
	if splitted[0] == "v" && len(splitted) != 4 {
		return false, false, point, nil
	} else if splitted[0] == "v" {
		for i := 1; i <= 3; i++ {
			if float, err := strconv.ParseFloat(splitted[i], 64); err == nil {
				cors = append(cors, float)
			} else {
				return false, false, point, err
			}
		}
		return true, true, cors, nil
	}
	if splitted[0] == "f" {
		for i := 1; i < len(splitted); i++ {
			if float, err := strconv.ParseFloat(splitted[i], 64); err == nil {
				cors = append(cors, float)
			} else {
				return false, false, point, err
			}
		}
		return true, false, cors, nil
	}
	if splitted[0] == "g"{
		return true, false, cors, nil
	}
	return false, false, point, finErr
}

func fanTriangulation(vertices []Tuple, point []float64) []interface{} {
	var tris []interface{}
	for i := 1; i < len(point)-1; i++ {
		tris = append(tris, *NewTriangle(&vertices[int(point[0])-1], &vertices[int(point[i])-1], &vertices[int(point[i+1])-1]))
	}
	return tris
}
