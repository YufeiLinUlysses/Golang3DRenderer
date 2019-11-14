package main

import (
	"feature"
	"fmt"
	//"method"
)

func main() {
	op := feature.NewOBJParser("../test/testFiles/originalFile/test.obj")
	op = op.ReadObj()
	//fmt.Println(op.OBJToGroup().Objects[1].(*feature.Group).Objects)
	fmt.Println(op.OBJToGroup().Objects[0].(*feature.Group).Objects[0].(feature.SmoothTriangle).Normal1)
	fmt.Println(op.OBJToGroup().Objects[0].(*feature.Group).Objects[0].(feature.SmoothTriangle).Normal2)
	fmt.Println(op.OBJToGroup().Objects[0].(*feature.Group).Objects[0].(feature.SmoothTriangle).Normal3)
	fmt.Println(op.OBJToGroup().Objects[0].(*feature.Group).Objects[1].(feature.SmoothTriangle).Normal1)
	fmt.Println(op.OBJToGroup().Objects[0].(*feature.Group).Objects[1].(feature.SmoothTriangle).Normal2)
	fmt.Println(op.OBJToGroup().Objects[0].(*feature.Group).Objects[1].(feature.SmoothTriangle).Normal3)
	// r := feature.NewRay(*feature.Point(-0.2, 0.3, -2), *feature.Vector(0, 0, 1))
	// p1 := feature.Point(0, 1, 0)
	// p2 := feature.Point(-1, 0, 0)
	// p3 := feature.Point(1, 0, 0)
	// n1 := feature.Vector(0, 1, 0)
	// n2 := feature.Vector(-1, 0, 0)
	// n3 := feature.Vector(1, 0, 0)
	// stri := feature.NewSmoothTriangle(*p1, *p2, *p3, *n1, *n2, *n3)
	// _, ans, _ := stri.IntersectWithRay(r)
	// fmt.Println(ans[0].U)
	// fmt.Println(ans[0].V)
	// fmt.Println(stri.NormalAt(feature.Point(0, 0, 0),ans[0]))
	// temp := ans[0].PrepareComputation(r,ans)
	// fmt.Println(temp.Normal)
}
