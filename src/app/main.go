package main

import (
	"feature"
	"fmt"
	"math"
	//"method"
)

func main() {
	// dir, _ := feature.Vector(0, 1, 1).Normalize()
	// r := feature.NewRay(*feature.Point(0, 0, -0.25), dir)
	// c := feature.NewCone()
	// c.Max = 0.5
	// c.Min = -0.5
	// c.Closed = true
	// fmt.Println(c.IntersectWithRay(r))
	//r := feature.NewRay(*feature.Point(10, 0, -10), *feature.Vector(0, 0, 1))
	g1 := feature.NewGroup()
	g1.Transform = feature.RotationY(math.Pi / 2)
	g2 := feature.NewGroup()
	g2.Transform = feature.Scale(1, 2, 3)
	g1.AddChild(g2)
	s := feature.NewSphere()
	s.Transform = feature.Translate(5, 0, 0)
	g2.AddChild(s)
	fmt.Println(s.NormalAt(feature.Point(1.7321, 1.1547, -5.5774)))
	//method.SixthImage("../../output/test7")
}
