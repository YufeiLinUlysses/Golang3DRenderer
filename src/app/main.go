package main

import (
	"feature"
	"fmt"
	//"method"
)

func main() {
	var items []feature.Intersection
	r := feature.NewRay(*feature.Point(0, 0, -4), *feature.Vector(0, 0, 1))

	a := feature.NewSphere()
	a.Transform = feature.Scale(2, 2, 2)
	a.Mat.Refractivity = 1.5

	b := feature.NewSphere()
	b.Transform = feature.Translate(0, 0, -0.25)
	b.Mat.Refractivity = 2.0

	c := feature.NewSphere()
	c.Transform = feature.Translate(0, 0, 0.25)
	c.Mat.Refractivity = 2.5

	i1 := feature.NewIntersection(2, *r, a)
	i2 := feature.NewIntersection(2.75, *r, b)
	i3 := feature.NewIntersection(3.25, *r, c)
	i4 := feature.NewIntersection(4.75, *r, b)
	i5 := feature.NewIntersection(5.25, *r, c)
	i6 := feature.NewIntersection(6, *r, a)

	items = append(items, *i1, *i2, *i3, *i4, *i5, *i6)
	for i := range items {
		comps := items[i].PrepareComputation(r, items)
		fmt.Println(comps.Refract1)
		fmt.Println(comps.Refract2)
		fmt.Println()
	}
}
