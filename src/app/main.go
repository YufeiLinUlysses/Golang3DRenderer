package main

import (
	"feature"
	//"fmt"
	//"method"
)

func main() {
	op := feature.NewOBJParser("../test/testFiles/originalFile/test.obj")
	op.ReadObj()
	//method.SixthImage("../../output/test7")
}
