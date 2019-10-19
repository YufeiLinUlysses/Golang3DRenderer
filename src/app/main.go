package main

import (
	"feature"
	"fmt"
	"math"
	//"method"
)

func main() {
	var light feature.Light
	var lights []feature.Light
	light = light.PointLight(*feature.Point(0, 0, 0), *feature.NewColor(1, 1, 1))
	lights = append(lights, light)
	w := feature.DefaultWorld()
	shape := feature.NewSphere()
	shape.Mat.Transparency = 1.0
	shape.Mat.Refractivity = 1.5
	shape.Transform = feature.Translate(0, -1, 0)
	w.Objects = append(w.Objects, shape)
	r := feature.NewRay(*feature.Point(0, 0, -3), *feature.Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	i := feature.NewIntersection(math.Sqrt(2), *r, shape)
	comps := i.PrepareComputation(r)
	color := w.ReflectedColor(comps, 0)
	fmt.Println(color)
}
