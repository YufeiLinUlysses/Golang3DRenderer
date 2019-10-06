package main

import (
	"feature"
	"fmt"
	"math"
	//"method"
)

func main() {
	//method.SecondImage("../../output/test2")
	m := feature.NewMaterial()
	position := feature.Point(0, 0, 0)
	eyev := feature.Vector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normal := feature.Vector(0, 0, -1)
	light := feature.NewLight()
	*light = light.PointLight(*feature.Point(0, 10, -10), *feature.NewColor(1, 1, 1))
	//s := feature.NewSphere()
	//s.Material.Ambient = 1
	// matrix := feature.Shearing(1, 0, 0, 0, 0, 0)
	// ans, _ := matrix.MultiplyTuple(point)
	fmt.Println(m.Lighting(*light, position, normal, eyev))
}
