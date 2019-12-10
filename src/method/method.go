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
				color := DiffuseLight(ans[0].Position, l, ray, red, *s)
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
				hitPoint := ray.Position(ans[0].Position)
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

//FifthImage creates the fifth image
func FifthImage(fileName string) {

	//Floor instance
	floor := feature.NewPlane()
	//floor.Transform = feature.Scale(10, 0.01, 10)
	floor.Mat.Col = *feature.NewColor(1, 0.9, 0.9)
	floor.Mat.Specular = 0

	//Left wall instance
	leftWall := feature.NewPlane()
	trans := feature.Translate(0, 0, 2.5)
	rotX := feature.RotationX(math.Pi / 2)
	rotY := feature.RotationY(-math.Pi / 3)
	m := rotX
	m, _ = rotY.Multiply(m)
	m, _ = trans.Multiply(m)
	leftWall.Transform = m
	leftWall.Mat = floor.Mat

	//Right wall instance
	rightWall := feature.NewPlane()
	transR := feature.Translate(0, 0, 4.7)
	rotYR := feature.RotationY(math.Pi / 3)
	rotX = feature.RotationX(math.Pi / 2)
	mR, _ := rotYR.Multiply(rotX)
	mR, _ = transR.Multiply(mR)
	rightWall.Transform = mR
	rightWall.Mat = floor.Mat

	//Middle wall instance
	middleWall := feature.NewPlane()
	transMid := feature.Translate(0, 0, 0.4)
	rotXMid := feature.RotationX(math.Pi / 2)
	mMid := rotXMid
	mMid, _ = transMid.Multiply(mMid)
	middleWall.Transform = mMid
	middleWall.Mat = floor.Mat

	//Right instance
	right := feature.NewSphere()
	right.Transform, _ = feature.Translate(1.5, 0.4, -1.5).Multiply(feature.Scale(0.5, 0.5, 0.5))
	right.Mat.Col = *feature.NewColor(1, 0.2, 1)
	right.Mat.Diffuse = 0.7
	right.Mat.Specular = 0.3

	//Left instance
	left := feature.NewSphere()
	left.Transform, _ = feature.Translate(-0.8, 0.45, 0.2).Multiply(feature.Scale(0.5, 0.5, 0.5))
	left.Mat.Col = *feature.NewColor(1, 0, 0)
	left.Mat.Diffuse = 0.7
	left.Mat.Specular = 0.3

	//Camera instance
	cam := feature.NewCamera(100, 65, math.Pi/3)
	cam.Transform = feature.ViewTransformation(*feature.Point(0, 1.5, -5), *feature.Point(0, 1, 0), *feature.Vector(0, 1, 0))

	//World instance
	var light feature.Light
	var lights []feature.Light
	var objects []interface{}

	//Add light source
	lights = append(lights, light.PointLight(*feature.Point(5.5, 20, -5), *feature.NewColor(1, 1, 1)))

	//Add all objects in
	objects = append(objects, middleWall, rightWall, leftWall, left, right, floor)
	w := feature.NewWorld(lights, objects)

	//Canv that draw the final product
	canv := cam.Render(*w)
	canv.CanvasToPPM(fileName)
}

//SixthImage creates the sixth image
func SixthImage(fileName string) {
	cube := feature.NewCube()
	cube.Mat.Col = *feature.NewColor(0.5, 1, 0.1)

	right := feature.NewSphere()
	right.Transform, _ = feature.Translate(1.5, 0.5, -0.5).Multiply(feature.Scale(0.5, 0.5, 0.5))
	right.Mat.Col = *feature.NewColor(0.5, 1, 0.1)
	right.Mat.Diffuse = 0.7
	right.Mat.Specular = 0.3

	cam := feature.NewCamera(100, 100, math.Pi/3)
	cam.Transform = feature.ViewTransformation(*feature.Point(0, 1.5, -5), *feature.Point(0, 1, 0), *feature.Vector(0, 1, 0))

	var light feature.Light
	var lights []feature.Light
	var objects []interface{}

	//Add light source
	lights = append(lights, light.PointLight(*feature.Point(-10, 10, -10), *feature.NewColor(1, 1, 1)))

	//Add all objects in
	objects = append(objects, cube, right)
	w := feature.NewWorld(lights, objects)

	//Canv that draw the final product
	canv := cam.Render(*w)
	canv.CanvasToPPM(fileName)
}

//SeventhImage creates the seventh image
func SeventhImage(fileName string) {
	torus := feature.NewTorus(0.4, 0.2)
	torus.Mat.Col = *feature.NewColor(0.5, 1, 0.1)
	//torus.Transform = feature.RotationX(math.Pi / 5)
	torus.Transform, _ = torus.Transform.Multiply(feature.RotationY(math.Pi / 3))
	torus.Transform, _ = torus.Transform.Multiply(feature.RotationZ(math.Pi / 2))
	cam := feature.NewCamera(1000, 1000, math.Pi/3)
	cam.Transform = feature.ViewTransformation(*feature.Point(0, 1.5, -5), *feature.Point(0, 1, 0), *feature.Vector(0.55, 0.85, 5.5))

	var light feature.Light
	var lights []feature.Light
	var objects []interface{}

	//Add light source
	lights = append(lights, light.PointLight(*feature.Point(-10, 10, -10), *feature.NewColor(1, 1, 1)))

	//Add all objects in
	objects = append(objects, torus)
	w := feature.NewWorld(lights, objects)

	//Canv that draw the final product
	canv := cam.Render(*w)
	canv.CanvasToPPM(fileName)
}

//EighthImage creates the eighth image
func EighthImage(fileName string) {
	//Floor instance
	floor := feature.NewPlane()
	floor.Mat.Col = *feature.NewColor(1, 0.9, 0.9)
	floor.Mat.Specular = 0

	//Left wall instance
	leftWall := feature.NewPlane()
	trans := feature.Translate(0, 0, 2.5)
	rotX := feature.RotationX(math.Pi / 2)
	rotY := feature.RotationY(-math.Pi / 3)
	m := rotX
	m, _ = rotY.Multiply(m)
	m, _ = trans.Multiply(m)
	leftWall.Transform = m
	leftWall.Mat = floor.Mat
	leftWall.Mat.Pat = *feature.NewPattern(*feature.NewColor(1, 0, 0), *feature.NewColor(0, 0.2, 0.78))
	leftWall.Mat.HasPattern = true
	leftWall.Mat.PatternType = "gradient"

	//Right wall instance
	rightWall := feature.NewPlane()
	transR := feature.Translate(0, 0, 4.7)
	rotYR := feature.RotationY(math.Pi / 3)
	rotX = feature.RotationX(math.Pi / 2)
	mR, _ := rotYR.Multiply(rotX)
	mR, _ = transR.Multiply(mR)
	rightWall.Transform = mR
	rightWall.Mat = floor.Mat
	rightWall.Mat.Pat = *feature.NewPattern(*feature.NewColor(1, 0.5, 0), *feature.NewColor(0, 0.747, 0.8))
	rightWall.Mat.HasPattern = true
	rightWall.Mat.PatternType = "ring"

	//Middle wall instance
	middleWall := feature.NewPlane()
	transMid := feature.Translate(0, 0, 0.4)
	rotXMid := feature.RotationX(math.Pi / 2)
	mMid := rotXMid
	mMid, _ = transMid.Multiply(mMid)
	middleWall.Transform = mMid
	middleWall.Mat = floor.Mat
	middleWall.Transform = mR
	middleWall.Mat = floor.Mat
	middleWall.Mat.Pat = *feature.NewPattern(*feature.NewColor(0.4, 0.2, 0), *feature.NewColor(0, 0.3, 0.8))
	middleWall.Mat.HasPattern = true
	middleWall.Mat.PatternType = "stripe"

	//Right instance
	right := feature.NewSphere()
	right.Transform, _ = feature.Translate(1.5, 0.4, -1.5).Multiply(feature.Scale(0.5, 0.5, 0.5))
	right.Mat.Col = *feature.NewColor(1, 0.2, 1)
	right.Mat.Diffuse = 0.7
	right.Mat.Specular = 0.3

	//Left instance
	left := feature.NewSphere()
	left.Transform, _ = feature.Translate(-0.8, 0.45, 0.2).Multiply(feature.Scale(0.5, 0.5, 0.5))
	left.Mat.Col = *feature.NewColor(1, 0, 0)
	left.Mat.Diffuse = 0.7
	left.Mat.Specular = 0.3
	left.Mat.Refractivity = 0.5
	left.Mat.Reflectivity = 0.8

	//Camera instance
	cam := feature.NewCamera(1000, 650, math.Pi/3)
	cam.Transform = feature.ViewTransformation(*feature.Point(0, 1.5, -5), *feature.Point(0, 1, 0), *feature.Vector(0, 1, 0))

	//World instance
	var light feature.Light
	var lights []feature.Light
	var objects []interface{}

	//Add light source
	lights = append(lights, light.PointLight(*feature.Point(5.5, 20, -5), *feature.NewColor(0.01, 0.21, 0.334)))

	//Add all objects in
	objects = append(objects, middleWall, rightWall, leftWall, left, right, floor)
	w := feature.NewWorld(lights, objects)

	//Canv that draw the final product
	canv := cam.Render(*w)
	canv.CanvasToPPM(fileName)
}

//NinthImage creates the ninth image
func NinthImage(fileName string) {
	//Floor instance
	floor := feature.NewPlane()
	floor.Mat.Col = *feature.NewColor(1, 0.9, 0.9)
	floor.Mat.Specular = 0

	//Left wall instance
	leftWall := feature.NewPlane()
	trans := feature.Translate(0, 0, 2.5)
	rotX := feature.RotationX(math.Pi / 2)
	rotY := feature.RotationY(-math.Pi / 3)
	m := rotX
	m, _ = rotY.Multiply(m)
	m, _ = trans.Multiply(m)
	leftWall.Transform = m
	leftWall.Mat = floor.Mat
	leftWall.Mat.Pat = *feature.NewPattern(*feature.NewColor(1, 0, 0), *feature.NewColor(0, 0.2, 0.78))
	leftWall.Mat.HasPattern = true
	leftWall.Mat.PatternType = "gradient"

	//Right wall instance
	rightWall := feature.NewPlane()
	transR := feature.Translate(0, 0, 4.7)
	rotYR := feature.RotationY(math.Pi / 3)
	rotX = feature.RotationX(math.Pi / 2)
	mR, _ := rotYR.Multiply(rotX)
	mR, _ = transR.Multiply(mR)
	rightWall.Transform = mR
	rightWall.Mat = floor.Mat
	rightWall.Mat.Pat = *feature.NewPattern(*feature.NewColor(1, 0.5, 0), *feature.NewColor(0, 0.747, 0.8))
	rightWall.Mat.HasPattern = true
	rightWall.Mat.PatternType = "ring"

	//Middle wall instance
	middleWall := feature.NewPlane()
	transMid := feature.Translate(0, 0, 0.4)
	rotXMid := feature.RotationX(math.Pi / 2)
	mMid := rotXMid
	mMid, _ = transMid.Multiply(mMid)
	middleWall.Transform = mMid
	middleWall.Mat = floor.Mat
	middleWall.Transform = mR
	middleWall.Mat = floor.Mat
	middleWall.Mat.Pat = *feature.NewPattern(*feature.NewColor(0.4, 0.2, 0), *feature.NewColor(0, 0.3, 0.8))
	middleWall.Mat.HasPattern = true
	middleWall.Mat.PatternType = "stripe"

	//Right Sphere instance
	right := feature.NewSphere()
	right.Transform, _ = feature.Translate(1.5, 0.4, -1.5).Multiply(feature.Scale(0.5, 0.5, 0.5))
	right.Mat.Col = *feature.NewColor(1, 0.2, 1)
	right.Mat.Diffuse = 0.7
	right.Mat.Specular = 0.3

	//Right Cube instance
	rcube := feature.NewCube()
	rcube.Transform, _ = feature.Translate(1.5, 0.4, -1.5).Multiply(feature.Scale(0.5, 0.5, 0.5))
	rcube.Transform, _ = feature.RotationX(math.Pi / 5).Multiply(rcube.Transform)
	rcube.Mat.Col = *feature.NewColor(1, 0.2, 1)
	rcube.Mat.Diffuse = 0.7
	rcube.Mat.Specular = 0.3
	rcube.Mat.Pat = *feature.NewPattern(*feature.NewColor(0.8, 0.2, 0), *feature.NewColor(0, 0.3, 0.8))
	rcube.Mat.HasPattern = true
	rcube.Mat.PatternType = "ring"
	rcube.Mat.Transparency = 0.9

	//Left instance
	left := feature.NewSphere()
	left.Transform, _ = feature.Translate(-0.8, 0.45, 0.2).Multiply(feature.Scale(0.5, 0.5, 0.5))
	left.Mat.Col = *feature.NewColor(1, 0, 0)
	left.Mat.Diffuse = 0.7
	left.Mat.Specular = 0.3
	left.Mat.Refractivity = 0.5
	left.Mat.Reflectivity = 0.8

	//Left Cube instance
	lcube := feature.NewCube()
	lcube.Transform, _ = feature.Translate(-0.3, 0, 2).Multiply(feature.Scale(0.2, 0.2, 0.2))
	lcube.Transform,_ = feature.RotationY(math.Pi/12).Multiply(lcube.Transform)
	lcube.Mat.Col = *feature.NewColor(1, 0.2, 1)
	lcube.Mat.Diffuse = 0.7
	lcube.Mat.Specular = 0.56
	lcube.Mat.Pat = *feature.NewPattern(*feature.NewColor(0, 0.2, 0), *feature.NewColor(0, 0.3, 0.8))
	lcube.Mat.HasPattern = true
	lcube.Mat.PatternType = "gradient"
	lcube.Mat.Transparency = 0.3
	lcube.Mat.Reflectivity = 1

	//Left Cylinder instance
	lcylinder := feature.NewCylinder()
	lcylinder.Max = 3
	lcylinder.Min = 1
	lcylinder.Transform, _ = feature.Translate(0.2, -0.6, 1.5).Multiply(feature.Scale(0.8, 0.8, 0.8))
	lcylinder.Mat.Col = *feature.NewColor(1, 0.2, 1)
	lcylinder.Mat.Diffuse = 0.7
	lcylinder.Mat.Specular = 0.56
	lcylinder.Mat.Pat = *feature.NewPattern(*feature.NewColor(0, 1, 0), *feature.NewColor(1, 0.2, 0.8))
	lcylinder.Mat.HasPattern = true
	lcylinder.Mat.PatternType = "stripe"
	lcylinder.Mat.Transparency = 0.9
	lcylinder.Mat.Reflectivity = 0.5

	//Cone instance
	cone := feature.NewCone()
	cone.Max =2
	cone.Min =-1.5
	cone.Transform,_ =  feature.Translate(1, 0.4, 1).Multiply(feature.Scale(0.5, 0.5, 0.5))
	cone.Mat.Col = *feature.NewColor(1, 0, 0)
	cone.Mat.Diffuse = 0.7
	cone.Mat.Specular = 0.3

	//Torus instance
	torus := feature.NewTorus(0.4,0.2)
	torus.Transform, _ =  feature.Translate(3, 0, -1).Multiply(feature.Scale(0.5, 0.5, 0.5))
	torus.Transform,_=feature.RotationY(math.Pi/3).Multiply(torus.Transform)
	torus.Mat.Col = *feature.NewColor(0, 1, 1)
	torus.Mat.Diffuse = 0.9
	torus.Mat.Specular = 0.3
	torus.Mat.Reflectivity = 0.2
	//Camera instance
	cam := feature.NewCamera(2000, 1300, math.Pi/3)
	cam.Transform = feature.ViewTransformation(*feature.Point(0, 1.5, -5), *feature.Point(0, 1, 0), *feature.Vector(0, 2, 2))

	//World instance
	var light feature.Light
	var lights []feature.Light
	var objects []interface{}

	//Add light source
	light1 :=light.PointLight(*feature.Point(5.5, 20, -5), *feature.NewColor(1, 1, 0))
	light2 := light.PointLight(*feature.Point(1.5, 34, -10), *feature.NewColor(0, 0, 0))
	lights = append(lights, light1, light2)

	//Add all objects in
	objects = append(objects, middleWall, rightWall, leftWall, lcube, lcylinder, left, right, cone, rcube, floor, torus)
	w := feature.NewWorld(lights, objects)

	//Canv that draw the final product
	canv := cam.Render(*w)
	canv.CanvasToPPM(fileName)
}
