package method

import (
	"class"
)

//FirstImage creates the first image
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

//SecondImage creates the second image
func SecondImage() {
	var x, y float64
	canv := class.NewCanvas(100, 100)
	red := class.NewColor(1, 0, 0)
	s := class.NewSphere()
	l := class.NewLight()
	l.PointLight(*class.Point(1, 0, -5), *class.NewColor(0.5, 1, 0.3))
	for i, row := range canv.Canv {
		y = float64(2) - float64(i)/float64(25)
		for j := range row {
			x = float64(-2) + float64(j)/float64(25)
			ray := class.NewRay(x, y, -5, 0, 0, 1)
			_, ans1,_, intersect := s.IntersectWithRay(ray)
			if intersect {
				color := DiffuseLight(ans1, l, ray, red, *s)
				canv.WritePixel(i, j, color)
			}
		}
	}
	canv.CanvasToPPM("test2")
}

//DiffuseLight diffuses Light
func DiffuseLight(intersection float64, l *class.Light, ray *class.Ray, originalColor *class.Color, s class.Sphere) *class.Color {
	hitPoint := ray.Position(intersection)
	vectorToLight, _ := l.Position.Subtract(&hitPoint)
	unitVectorToLight, _ := vectorToLight.Normalize()
	normal := s.NormalAt(&hitPoint)
	lightIntensity, _ := normal.DotProduct(&unitVectorToLight)
	colorAtPoint := l.Intensity.ColorMultiply(originalColor)
	color := colorAtPoint.Multiply(lightIntensity)
	return &color
}
