package main

import (
	"feature"
	"fmt"
	//"method"
)

func main() {
	r := feature.NewRay(*feature.Point(0, 0, -5), *feature.Vector(0, 0, 1))
	c := feature.NewCylinder()
	fmt.Println(c.IntersectWithRay(r))
}
