package main

import (
	"class"
	"fmt"
	//"method"
)

func main() {
	//method.SecondImage("../../output/test2")
	r := class.NewRay(0, 0, -5, 0, 0, 1)
	matrix := class.Scale(2, 2, 2)
	s := class.NewSphere()
	s = s.SetTransform(matrix)
	fmt.Println(s.Transform)
	fmt.Println(s.IntersectWithRay(r))
}
