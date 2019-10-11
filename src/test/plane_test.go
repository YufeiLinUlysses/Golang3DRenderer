package test

import (
	"feature"
	"testing"
)

//TestPlane1 tests to see if the GetColor  function works for feature Color
func TestPlane1(t *testing.T) {
	tables := []struct {
		r     *feature.Ray
		ansED bool
		ans   float64
	}{
		{feature.NewRay(*feature.Point(0, 10, 0), *feature.Vector(0, 0, 1)), false, 0},
		{feature.NewRay(*feature.Point(0, 0, 0), *feature.Vector(0, 0, 1)), false, 0},
		{feature.NewRay(*feature.Point(0, 1, 0), *feature.Vector(0, -1, 0)), true, 1},
		{feature.NewRay(*feature.Point(0, -1, 0), *feature.Vector(0, 1, 0)), true, 1},
	}
	for _, table := range tables {
		p := feature.NewPlane()
		_, ans, ansED := p.IntersectWithRay(table.r)
		if ansED {
			if ans.T != table.ans {
				t.Errorf("Error Input")
			}
		} else {
			if ansED != table.ansED {
				t.Errorf("Error Input")
			}
		}
	}
}
