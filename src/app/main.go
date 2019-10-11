package main

import (
	"feature"
	"fmt"
	//"method"
)

func main() {
	var lights []feature.Light
	var objects []interface{}
	l := feature.NewLight()
	*l = l.PointLight(*feature.Point(0, 0, -10), *feature.NewColor(1, 1, 1))
	lights = append(lights, *l)
	s1 := feature.NewSphere()
	s2 := feature.NewSphere()
	s2.Transform = feature.Translate(0, 0, 10)
	objects = append(objects, s1, s2)
	w := feature.NewWorld(lights, objects)
	r := feature.NewRay(*feature.Point(0, 0, 5), *feature.Vector(0, 0, 1))
	i := feature.NewIntersection(4, *r, s2)
	comp := i.PrepareComputation(r)
	c := w.ShadeHit(comp)
	fmt.Println(c)
	// c.Transform = feature.ViewTransformation(*from, *to, *up)
	// image := c.Render(*w)
	// fmt.Println(image.PixelAt(5, 5))
	//method.ForthImage("../../output/test4")
}
