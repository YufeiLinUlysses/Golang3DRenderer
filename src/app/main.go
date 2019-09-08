package main

import (
	"class"
	"fmt"
)

func main() {
	cone := class.NewCanvas(2, 3)
	//ctwo := class.NewColor(0.9, 1, 0.1)
	//vtwo := class.NewTuple(2, 3, 4, 0)
	fmt.Println(cone.PixelAt(0, 1))
}
