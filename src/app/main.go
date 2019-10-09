package main

import (
	"feature"
	"fmt"
	//"method"
)

func main() {
	//t := &Test{A: 2}
	w := feature.DefaultWorld()
	r := feature.NewRay(*feature.Point(0, 0, -5), *feature.Vector(0, 0, 1))
	fmt.Println(w.IntersectWorld(r))
	//method.ThirdImage("../../output/test3")
}
