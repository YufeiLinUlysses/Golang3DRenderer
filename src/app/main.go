package main

import (
	"class"
	"fmt"
)

func main() {
	cone := class.NewCanvas(2, 3)
	ctwo := class.NewColor(0.9, 1, 0.1)
	// //vtwo := class.NewTuple(2, 3, 4, 0)
	cone.WritePixel(1, 2, ctwo)
	// fmt.Println(cone.PixelAt(0, 1))
	PPM := cone.CanvasToString()
	class.CanvasToPPM(PPM, "hahaha")
	fmt.Println(PPM)
	// rone := class.NewRay(0, 0, 5, 0, 0, 1)
	// sone := class.NewSphere()
	// fmt.Println(sone.IntersectWithRay(rone))
}
