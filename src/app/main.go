package main

import (
	"feature"
	"fmt"
	//"method"
)

func main() {
	op := feature.NewOBJParser("../test/testFiles/originalFile/test.obj")
	op = op.ReadObj()
	fmt.Println(len(op.OBJToGroup().Objects[1].(*feature.Group).Objects))
}
