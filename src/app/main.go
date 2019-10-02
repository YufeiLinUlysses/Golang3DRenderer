package main

import (
	"feature"
	"fmt"
	//"method"
)

func main() {
	//method.SecondImage("../../output/test2")
	point := feature.Point(2,3,4)
	matrix := feature.Shearing(1, 0, 0, 0, 0, 0)
	ans, _ := matrix.MultiplyTuple(point)
	fmt.Println(ans)
}
