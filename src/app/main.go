package main

import (
	"feature"
	"fmt"
	"math"
	//"method"
)

func main() {
	s := feature.NewPlane()
	r := feature.NewRay(*feature.Point(0, 1, -1), *feature.Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	i := feature.NewIntersection(math.Sqrt(2), *r, s)
	comp := i.PrepareComputation(r)
	fmt.Println(comp.Reflect)
	//method.FifthImage("../../output/PhysicsFinal1")
}
