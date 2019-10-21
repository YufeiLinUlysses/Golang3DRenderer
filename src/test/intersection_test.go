package test

import (
	"feature"
	"fmt"
	"testing"
)

//TestIntersection1 tests to see if the Hit function works for feature Intersection
func TestIntersection1(t *testing.T) {
	tables := []struct {
		t    []float64
		ans  float64
		ansB bool
	}{
		{[]float64{1, 2}, 1, true},
		{[]float64{-1, 1}, 1, true},
		{[]float64{-2, -1}, 0, false},
		{[]float64{5, 7, -3, 2}, 2, true},
	}
	for _, table := range tables {
		var inter []feature.Intersection
		w := feature.DefaultWorld()
		r := feature.NewRay(*feature.Point(0, 0, -5), *feature.Vector(0, 0, 1))
		for i := range table.t {
			inter = append(inter, *feature.NewIntersection(table.t[i], *r, w))
		}
		ans, ansB := feature.Hit(inter)
		if ansB {
			if ans.T != table.ans || ansB != table.ansB {
				t.Errorf("Error Input %v, %v", ans, ansB)
			}
		} else {
			if ansB != table.ansB {
				t.Errorf("Error Input %v, %v", ans, ansB)
			}
		}

	}
}

//TestIntersection2 tests to see if the PrepareComputation function works for feature Intersection by giving the refractive index
func TestIntersection2(t *testing.T) {
	tables := []struct {
		num      int
		refract1 float64
		refract2 float64
	}{
		{0, 1, 1.5},
		{1, 1.5, 2},
		{2, 2, 2.5},
		{3, 2.5, 2.5},
		{4, 2.5, 1.5},
		{5, 1.5, 1},
	}
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
	for _, table := range tables {
		comps := items[table.num].PrepareComputation(r, items)
		if comps.Refract1 != table.refract1 || comps.Refract2 != table.refract2 {
			fmt.Println(comps.Refract1)
			fmt.Println(comps.Refract2)
			fmt.Println(table.refract1)
			fmt.Println(table.refract2)
			t.Errorf("Error Input")
		}
	}
}
