package main

import (
	"class"
	"fmt"
)

func main() {
	vone := class.NewTuple(1, 2, 3, 0)
	//vtwo := class.NewTuple(2, 3, 4, 0)
	fmt.Println(vone.Normalize())
}
