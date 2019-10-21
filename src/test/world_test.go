package test

import (
	"feature"
	"math"
	"testing"
)

//TestWorld1 tests to see if the ColorAt function works for feature World
func TestWorld1(t *testing.T) {
	tables := []struct {
		r    *feature.Ray
		ansC *feature.Color
	}{
		{feature.NewRay(*feature.Point(0, 0, -5), *feature.Vector(0, 1, 0)), feature.NewColor(0, 0, 0)},
		{feature.NewRay(*feature.Point(0, 0, -5), *feature.Vector(0, 0, 1)), feature.NewColor(0.38066, 0.47583, 0.2855)},
	}
	for _, table := range tables {
		w := feature.DefaultWorld()
		c := w.ColorAt(table.r,5)
		errorAllowance := 0.00001
		if math.Abs(c.R-table.ansC.R) > errorAllowance || math.Abs(c.G-table.ansC.G) > errorAllowance || math.Abs(c.B-table.ansC.B) > errorAllowance {
			t.Errorf("Error Input %v, %v", c,table.ansC)
		}
	}
}
