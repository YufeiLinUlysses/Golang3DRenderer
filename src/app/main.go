package main

import (
	"class"
	"fmt"
)

func main() {
	cone := class.NewColor(1, 0.2, 0.4)
	ctwo := class.NewColor(0.9, 1, 0.1)
	//vtwo := class.NewTuple(2, 3, 4, 0)
	fmt.Println(cone.ColorMultiply(ctwo))
}
