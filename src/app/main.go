package main

import "class"

func main() {
	// m := class.NewLight()
	// fmt.Println(m.PointLight(*class.Point(0,0,0),*class.NewColor(1,1,1)))
	FirstImage()
}

func FirstImage() {
	var x, y float64
	canv := class.NewCanvas(100, 100)
	red := class.NewColor(1, 0, 0)
	s := class.NewSphere()
	for i, row := range canv.Canv {
		y = float64(2) - float64(i)/float64(25)
		for j := range row {
			x = float64(-2) + float64(j)/float64(25)
			ray := class.NewRay(x, y, -5, 0, 0, 1)
			_, _, _, intersect := s.IntersectWithRay(ray)
			if intersect {
				canv.WritePixel(i, j, red)
			}
		}
	}
	canv.CanvasToPPM("test1")
}
