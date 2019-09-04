package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y, z, w float64
}

func (v Vertex) isPoint() bool {
	if v.w == 1.0 {
		return true
	}
	return false
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := Vertex{4.3, -4.2, 3.1, 1.0}
	fmt.Println(v.Abs())
}
