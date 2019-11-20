package test

import (
	"feature"
	"math"
	"testing"
)

//TestWorld1 tests to see if the NewWorld function works for feature World
func TestWorld1(t *testing.T) {
	tables := []struct {
		ansL1, ansL2 int
	}{
		{0, 0},
	}
	for _, table := range tables {
		testLight := make([]feature.Light, 0)
		testShpae := make([]interface{}, 0)
		w := feature.NewWorld(testLight, testShpae)
		if len(w.Objects) != table.ansL1 || len(w.Lights) != table.ansL2 {
			t.Errorf("Error")
		}
	}
}

//TestWorld2 tests to see if the ColorAt function works for feature World
func TestWorld2(t *testing.T) {
	tables := []struct {
		r    *feature.Ray
		ansC *feature.Color
	}{
		{feature.NewRay(*feature.Point(0, 0, -5), *feature.Vector(0, 1, 0)), feature.NewColor(0, 0, 0)},
		{feature.NewRay(*feature.Point(0, 0, -5), *feature.Vector(0, 0, 1)), feature.NewColor(0.38066, 0.47583, 0.2855)},
	}
	for _, table := range tables {
		w := feature.DefaultWorld()
		c := w.ColorAt(table.r, 5)
		errorAllowance := 0.00001
		if math.Abs(c.R-table.ansC.R) > errorAllowance || math.Abs(c.G-table.ansC.G) > errorAllowance || math.Abs(c.B-table.ansC.B) > errorAllowance {
			t.Errorf("Error Input %v, %v", c, table.ansC)
		}
	}
}
