package main

import (
	"feature"
	"fmt"
	//"method"
)

func main() {
	from := feature.Point(0, 0, 0)
	to := feature.Point(0, 0, 1)
	up := feature.Vector(0, 1, 0)
	fmt.Println(up.Normalize())
	m := feature.ViewTransformation(*from, *to, *up)
	fmt.Println(m)
	fmt.Println(feature.Translate(0, 0, -8))
	//method.ThirdImage("../../output/test3")
}
