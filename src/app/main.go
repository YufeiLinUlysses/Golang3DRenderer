package main

import (
	"feature"
	"fmt"
	//"method"
)

func main() {
	r := feature.NewRay(*feature.Point(0, 0, -5), *feature.Vector(0, 0, 1))
	s1 := feature.NewSphere()
	s2 := feature.NewSphere()
	s2.Transform = feature.Translate(0, 0, 0.5)
	csg1 := feature.NewCSG("union", s1, s2)
	fmt.Println(csg1.IntersectWithRay(r))
}
