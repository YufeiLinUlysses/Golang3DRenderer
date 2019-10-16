package main

import (
	"feature"
	"fmt"
	//"method"
)

func main() {
	var light feature.Light
	s := feature.NewSphere()
	s.Transform = feature.Scale(2, 2, 2)
	light = light.PointLight(*feature.Point(0, 0, -10), *feature.NewColor(1, 1, 1))
	white := feature.NewColor(1, 1, 1)
	black := feature.NewColor(0, 0, 0)
	m := feature.NewMaterial()
	m.Pat = *feature.NewPattern(*white, *black)
	m.HasPattern = true
	m.Ambient = 1
	m.Diffuse = 0
	m.Specular = 0
	m.Pat.Transform = feature.Translate(0.5, 0, 0)
	comp := &feature.Computations{
		Point:  *feature.Point(0.9, 0, 0),
		Eye:    *feature.Vector(0, 0, -1),
		Normal: *feature.Vector(0, 0, -1),
		Shape:  *s,
	}
	fmt.Println(comp.Shape.(feature.Sphere).Transform)
	fmt.Println(m.Lighting(light, *comp, false))
	fmt.Println(m.Pat.StripeAt(*feature.Point(2.5, 0, 0), *s.Transform))
	//method.FifthImage("../../output/PhysicsFinal1")
}
