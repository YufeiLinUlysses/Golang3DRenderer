package method

import (
	"feature"
	"math"
)

//FirstImage creates the first image
func FirstImage() {
	var x, y float64
	canv := feature.NewCanvas(100, 100)
	red := feature.NewColor(1, 0, 0)
	s := feature.NewSphere()
	for i, row := range canv.Canv {
		y = float64(2) - float64(i)/float64(25)
		for j := range row {
			x = float64(-2) + float64(j)/float64(25)
			ray := feature.NewRay(*feature.Point(x, y, -5), *feature.Vector(0, 0, 1))
			_, _, intersect := s.IntersectWithRay(ray)
			if intersect {
				canv.WritePixel(i, j, red)
			}
		}
	}
	canv.CanvasToPPM("../../output/test1")
}

//SecondImage creates the second image
func SecondImage(fileName string) {
	var x, y float64
	canv := feature.NewCanvas(100, 100)
	red := feature.NewColor(1, 0, 0)
	s := feature.NewSphere()
	//Add Light in here
	l := feature.NewLight()
	l.PointLight(*feature.Point(1, 0, -5), *feature.NewColor(0.5, 1, 0.3))
	for i, row := range canv.Canv {
		y = float64(2) - float64(i)/float64(25)
		for j := range row {
			x = float64(-2) + float64(j)/float64(25)
			ray := feature.NewRay(*feature.Point(x, y, -5), *feature.Vector(0, 0, 1))
			_, ans, intersect := s.IntersectWithRay(ray)
			if intersect {
				color := DiffuseLight(ans[0].T, l, ray, red, *s)
				canv.WritePixel(i, j, color)
			}
		}
	}
	canv.CanvasToPPM(fileName)
}

//DiffuseLight diffuses Light
func DiffuseLight(intersection float64, l *feature.Light, ray *feature.Ray, originalColor *feature.Color, s feature.Sphere) *feature.Color {
	hitPoint := ray.Position(intersection)
	vectorToLight, _ := l.Position.Subtract(&hitPoint)
	unitVectorToLight, _ := vectorToLight.Normalize()
	normal := s.NormalAt(&hitPoint)
	lightIntensity, _ := normal.DotProduct(&unitVectorToLight)
	colorAtPoint := l.Intensity.ColorMultiply(originalColor)
	color := colorAtPoint.Multiply(lightIntensity)
	return &color
}

//ThirdImage creates the third image
func ThirdImage(fileName string) {
	var x, y float64

	canv := feature.NewCanvas(100, 100)

	//Sphere
	s := feature.NewSphere()
	s.Mat.Col = *feature.NewColor(1, 0.2, 1)

	//Add Light in here
	l := feature.NewLight()
	*l = l.PointLight(*feature.Point(-10, 10, -10), *feature.NewColor(1, 1, 1))
	for i, row := range canv.Canv {
		y = float64(2) - float64(i)/float64(25)
		for j := range row {
			x = float64(-2) + float64(j)/float64(25)
			ray := feature.NewRay(*feature.Point(x, y, -5), *feature.Vector(0, 0, 1))
			_, ans, intersect := s.IntersectWithRay(ray)
			eye := ray.Direction.Multiply(-1)
			if intersect {
				hitPoint := ray.Position(ans[0].T)
				normal := s.NormalAt(&hitPoint)
				var comp feature.Computations
				comp.Eye = eye
				comp.Normal = normal
				comp.Point = hitPoint
				color := s.Mat.Lighting(*l, comp, false)
				canv.WritePixel(i, j, &color)
			}
		}
	}
	canv.CanvasToPPM(fileName)
}

//ForthImage creates the forth image
func ForthImage(fileName string) {

	//Floor instance
	floor := feature.NewSphere()
	floor.Transform = feature.Scale(10, 0.01, 10)
	floor.Mat.Col = *feature.NewColor(1, 0.9, 0.9)
	floor.Mat.Specular = 0

	//Left wall instance
	leftWall := feature.NewSphere()
	trans := feature.Translate(0, 0, 5)
	rotY := feature.RotationY(-math.Pi / 4)
	rotX := feature.RotationX(math.Pi / 2)
	scal := feature.Scale(10, 0.01, 10)
	m, _ := rotX.Multiply(scal)
	m, _ = rotY.Multiply(m)
	m, _ = trans.Multiply(m)
	leftWall.Transform = m
	leftWall.Mat = floor.Mat

	//Right wall instance
	rightWall := feature.NewSphere()
	trans = feature.Translate(0, 0, 5)
	rotY = feature.RotationY(math.Pi / 4)
	rotX = feature.RotationX(math.Pi / 2)
	scal = feature.Scale(10, 0.01, 10)
	m, _ = rotX.Multiply(scal)
	m, _ = rotY.Multiply(m)
	m, _ = trans.Multiply(m)
	rightWall.Transform = m
	rightWall.Mat = floor.Mat

	//Middle instance
	middle := feature.NewSphere()
	middle.Transform = feature.Translate(-0.5, 1, 0.5)
	middle.Mat.Col = *feature.NewColor(0.1, 1, 0.5)
	middle.Mat.Diffuse = 0.7
	middle.Mat.Specular = 0.3

	//Right instance
	right := feature.NewSphere()
	right.Transform, _ = feature.Translate(1.5, 0.5, -0.5).Multiply(feature.Scale(0.5, 0.5, 0.5))
	right.Mat.Col = *feature.NewColor(0.5, 1, 0.1)
	right.Mat.Diffuse = 0.7
	right.Mat.Specular = 0.3

	//Left instance
	left := feature.NewSphere()
	left.Transform, _ = feature.Translate(-1.5, 0.33, -0.75).Multiply(feature.Scale(0.33, 0.33, 0.33))
	left.Mat.Col = *feature.NewColor(1, 0.8, 0.1)
	left.Mat.Diffuse = 0.7
	left.Mat.Specular = 0.3

	//Camera instance
	cam := feature.NewCamera(100, 50, math.Pi/3)
	cam.Transform = feature.ViewTransformation(*feature.Point(0, 1.5, -5), *feature.Point(0, 1, 0), *feature.Vector(0, 1, 0))

	//World instance
	var lights []feature.Light
	var objects []interface{}
	var light feature.Light
	lights = append(lights, light.PointLight(*feature.Point(-10, 10, -10), *feature.NewColor(1, 1, 1)))
	objects = append(objects, floor, leftWall, rightWall, middle, right, left)
	w := feature.NewWorld(lights, objects)

	//Canv that draw the final product
	canv := cam.Render(*w)
	canv.CanvasToPPM(fileName)
}

//FifthImage creates the forth image
func FifthImage(fileName string) {

	//Floor instance
	floor := feature.NewPlane()
	//floor.Transform = feature.Scale(10, 0.01, 10)
	//floor.Mat.Col = *feature.NewColor(1, 0.9, 0.9)
	floor.Mat.Specular = 0

	//Left wall instance
	leftWall := feature.NewPlane()
	trans := feature.Translate(0, 0, 5)
	rotY := feature.RotationY(-math.Pi / 4)
	rotX := feature.RotationX(math.Pi / 2)
	scal := feature.Scale(10, 0.01, 10)
	m, _ := rotX.Multiply(scal)
	m, _ = rotY.Multiply(m)
	m, _ = trans.Multiply(m)
	leftWall.Transform = m
	leftWall.Mat = floor.Mat

	//Right wall instance
	rightWall := feature.NewPlane()
	trans = feature.Translate(0, 0, 5)
	rotY = feature.RotationY(math.Pi / 4)
	rotX = feature.RotationX(math.Pi / 2)
	m, _ = rotY.Multiply(rotX)
	m, _ = trans.Multiply(m)
	rightWall.Transform = m
	rightWall.Mat = floor.Mat

	//Middle instance
	middle := feature.NewSphere()
	middle.Transform = feature.Translate(-0.5, 1, 0.5)
	middle.Mat.Col = *feature.NewColor(0.1, 1, 0.5)
	middle.Mat.Diffuse = 0.7
	middle.Mat.Specular = 0.3

	//Right instance
	right := feature.NewSphere()
	right.Transform, _ = feature.Translate(1.5, 0.5, -0.5).Multiply(feature.Scale(0.5, 0.5, 0.5))
	right.Mat.Col = *feature.NewColor(0.5, 1, 0.1)
	right.Mat.Diffuse = 0.7
	right.Mat.Specular = 0.3

	//Left instance
	left := feature.NewSphere()
	left.Transform, _ = feature.Translate(-1.5, 0.33, -0.75).Multiply(feature.Scale(0.33, 0.33, 0.33))
	left.Mat.Col = *feature.NewColor(1, 0.8, 0.1)
	left.Mat.Diffuse = 0.7
	left.Mat.Specular = 0.3

	//Camera instance
	cam := feature.NewCamera(100, 50, math.Pi/3)
	cam.Transform = feature.ViewTransformation(*feature.Point(0, 1.5, -5), *feature.Point(0, 1, 0), *feature.Vector(0, 1, 0))

	//World instance
	var light feature.Light
	var lights []feature.Light
	var objects []interface{}

	//Add light source
	lights = append(lights, light.PointLight(*feature.Point(-10, 10, -10), *feature.NewColor(1, 1, 1)))

	//Add all objects in
	objects = append(objects, leftWall, rightWall, middle, left, right, floor)
	w := feature.NewWorld(lights, objects)

	//Canv that draw the final product
	canv := cam.Render(*w)
	canv.CanvasToPPM(fileName)
}
