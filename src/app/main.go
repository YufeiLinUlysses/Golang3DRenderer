package main

import (
	"feature"
	"fmt"
	//"method"
)

func main() {
	var light feature.Light
	s := feature.NewSphere()
	s.Transform = feature.Scale(1, 1, 1)
	fmt.Println(s.Transform)
	light = light.PointLight(*feature.Point(0, 0, -10), *feature.NewColor(1, 1, 1))
	white := feature.NewColor(1, 1, 1)
	black := feature.NewColor(0, 0, 0)
	m := feature.NewMaterial()
	m.Pat = *feature.NewPattern(*white, *black)
	m.HasPattern = true
	m.Ambient = 1
	m.Diffuse = 0
	m.Specular = 0
	m.PatternType = "ring"
	m.Pat.Transform = feature.Translate(0.5, 1, 1.5)
	comp := &feature.Computations{
		Point:  *feature.Point(2.5, 3, 3.5),
		Eye:    *feature.Vector(0, 0, -1),
		Normal: *feature.Vector(0, 0, -1),
		Shape:  s,
	}
	fmt.Println(m.Lighting(light, *comp, false))
	fmt.Println(m.Pat.PatternAt(*feature.Point(0, 0, 0), *s.Transform, m.PatternType))
	//method.FifthImage("../../output/PhysicsFinal1")
}
