package main

import (
	"feature"
	"fmt"
	//"method"
)

func main() {
	//method.SecondImage("../../output/test2")
	r := feature.NewRay(0, 0, -5, 0, 0, 1)
	matrix := feature.Scale(2, 2, 2)
	s := feature.NewSphere()
	s = s.SetTransform(matrix)
	fmt.Println(s.Transform)
	fmt.Println(s.IntersectWithRay(r))
}
