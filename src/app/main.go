package main

import (
	"feature"
	"fmt"
	//"method"
)

func main() {
	//m := feature.NewMaterial()
	w := feature.DefaultWorld()
	outer := w.Objects[0].(*feature.Sphere)
	inner := w.Objects[1].(*feature.Sphere)
	outer.Mat.Ambient = 1
	inner.Mat.Ambient = 1
	//w.Lights[0] = w.Lights[0].PointLight(*feature.Point(0, 0.25, 0), *feature.NewColor(1, 1, 1))
	r := feature.NewRay(*feature.Point(0, 0, 0.75), *feature.Vector(0, 0, -1))
	//i := feature.NewIntersection(4, *r, s)
	//comp := i.PrepareComputation(r)
	c := w.ColorAt(r)
	fmt.Println(c)
	fmt.Println(inner.Mat.Col)
	//method.ThirdImage("../../output/test3")
}
