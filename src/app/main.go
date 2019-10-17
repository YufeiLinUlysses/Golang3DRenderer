package main

import (
	"feature"
	//"method"
)

func main() {
	var light feature.Light
	var lights []feature.Light
	var objects []interface{}
	light = light.PointLight(*feature.Point(0, 0, 0), *feature.NewColor(1, 1, 1))
	lights = append(lights, light)
	lower := feature.NewPlane()
	lower.Mat.Reflective = 1
	lower.Transform = feature.Translate(0, -1, 0)
	upper := feature.NewPlane()
	upper.Mat.Reflective = 1
	upper.Transform = feature.Scale(0, 1, 0)
	objects = append(objects, lower, upper)
	w := feature.NewWorld(lights, objects)
	r := feature.NewRay(*feature.Point(0, 0, 0), *feature.Vector(0, 1, 0))
    w.ColorAt(r,0)
	// shape := feature.NewPlane()
	// shape.Mat.Reflective = 0.5
	// shape.Transform = feature.Translate(0, -1, 0)
	// w.Objects = append(w.Objects, shape)

	// i := feature.NewIntersection(math.Sqrt(2), *r, shape)
	// comp := i.PrepareComputation(r)
	// fmt.Println(w.ShadeHit(comp))
	//method.FifthImage("../../output/PhysicsFinal1")
}
