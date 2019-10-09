package test

import (
	"feature"
	"testing"
)

//TestColor1 tests to see if the Hit function works for feature Intersection
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
