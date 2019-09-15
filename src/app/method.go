package main

import (
	"class"
)

func FirstImage() {
	var x,y float64
	canv := class.NewCanvas(100, 100)
	red := class.NewColor(1,0,0)
	s := class.NewSphere()
	for i, row := range canv {
		y = 2-i/25 
		for j := range row {
			x=j/25-2
			ray := class.NewRay(x,y,-5,0,0,1)
			_,i1,i2,intersect := s.IntersectWithRay(ray)
			if intersect{
				canv.WritePixel(i,j,red)
			}
		}
	}
	canv.CanvasToPPM("test1")
}
