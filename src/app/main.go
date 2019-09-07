package main

import (
	"class"
	"fmt"
)

func main() {
	cone := class.NewColor(1.5, 2, 3)
	ctwo := class.NewColor(1, 2, 3)
	//vtwo := class.NewTuple(2, 3, 4, 0)
	fmt.Println(cone.Add(ctwo))
}
