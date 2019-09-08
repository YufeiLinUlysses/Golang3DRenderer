package main

import (
	"class"
	"fmt"
)

func main() {
	rone := class.NewRay(1, 0.2, 0.4, 1, 0, 0)
	//ctwo := class.NewColor(0.9, 1, 0.1)
	//vtwo := class.NewTuple(2, 3, 4, 0)
	fmt.Println(rone.Position(2))
}
