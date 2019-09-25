package main

import (
	"class"
	"fmt"
	//"fmt"
	//"method"
)

func main() {
	//method.SecondImage("../../output/test2")
	r := class.NewRay(1, 2, 3, 0, 1, 0)
	matrix := r.Scale(2,3,4)
	fmt.Println(matrix)
	fmt.Println(r.Transform(matrix))
}
