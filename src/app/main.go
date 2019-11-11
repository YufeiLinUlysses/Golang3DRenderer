package main

import (
	"feature"
	"fmt"
	//"method"
)

func main() {
	p1 := feature.Point(0, 1, 0)
	p2 := feature.Point(-1, 0, 0)
	p3 := feature.Point(1, 0, 0)
	tri := feature.NewTriangle(p1, p2, p3)
	fmt.Println(tri.Edge2)
	//method.SixthImage("../../output/test7")
}
