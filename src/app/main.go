package main

import (
	"class"
	"fmt"
	//"method"
)

func main() {
	//method.SecondImage("../../output/test2")
	r := class.NewRay(1, 2, 3, 0, 1, 0)
	matrix := class.Scale(2, 3, 4)
	fmt.Println(r.Transform(matrix))
}
