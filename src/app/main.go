package main

import (
	"feature"
	"fmt"
	//"method"
)

func main() {
	w := feature.DefaultWorld()
	//r := feature.NewRay(*feature.Point(0, 0, -5), *feature.Vector(0, 0, 1))
	fmt.Println(w.Objects[0])
	//method.ThirdImage("../../output/test3")
}
