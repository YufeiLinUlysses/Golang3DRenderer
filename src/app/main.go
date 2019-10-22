package main

import (
	"feature"
	"fmt"
	"math"
	//"method"
)

func main() {
	w := feature.DefaultWorld()
	r := feature.NewRay(*feature.Point(0, 0, -3), *feature.Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	floor := feature.NewPlane()
	floor.Transform = feature.Translate(0, -1, 0)
	floor.Mat.Reflectivity = 0.5
	floor.Mat.Transparency = 0.5
	floor.Mat.Refractivity = 1.5
	ball := feature.NewSphere()
	ball.Mat.Col = *feature.NewColor(1, 0, 0)
	ball.Mat.Ambient = 0.5
	ball.Transform = feature.Translate(0, -3.5, -0.5)
	w.Objects = append(w.Objects, floor, ball)
	i1 := feature.NewIntersection(math.Sqrt(2), *r, floor)
	xs := make([]feature.Intersection, 1)
	xs[0] = *i1
	comps := xs[0].PrepareComputation(r, xs)
	fmt.Println(w.ShadeHit(comps, 5))
}
