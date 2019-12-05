package main

import (
	"method"
)

func main() {
	// r := feature.NewRay(*feature.Point(0, 0, -5), *feature.Vector(0, 0, 1))
	// s1 := feature.NewSphere()
	// s2 := feature.NewSphere()
	// s2.Transform = feature.Translate(0, 0, 0.5)
	// csg1 := feature.NewCSG("union", s1, s2)
	// fmt.Println(csg1.IntersectWithRay(r))
	// coef := make([]float64, 5)
	// coef[0] = 5
	// coef[1] = 10
	// coef[2] = 10
	// coef[3] = 5
	// coef[4] = 1
	// fmt.Println(feature.SolveQuartic(coef))
	method.SeventhImage("../../output/test12")
}
