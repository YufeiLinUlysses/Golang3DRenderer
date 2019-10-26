package test

import (
	"feature"
	"testing"
)

//TestCylinder1 tests to see if the IntersectWithRay function works for feature Cylinder
func TestCylinder1(t *testing.T) {
	tables := []struct {
		r          feature.Ray
		count      int
		ans1, ans2 float64
	}{
		{*feature.NewRay(*feature.Point(1, 0, -5), *feature.Vector(0, 0, 1)), 2, 5, 5},
		{*feature.NewRay(*feature.Point(0, 0, -5), *feature.Vector(0, 0, 1)), 2, 4, 6},
		{*feature.NewRay(*feature.Point(0.5, 0, -5), *feature.Vector(0.1, 1, 1)), 2, 6.80798191702732, 7.088723439378861},
		//Misses a cylinder
		{*feature.NewRay(*feature.Point(1, 0, 0), *feature.Vector(0, 1, 0)), 0, 0, 0},
		{*feature.NewRay(*feature.Point(0, 0, 0), *feature.Vector(0, 1, 0)), 0, 0, 0},
		{*feature.NewRay(*feature.Point(0, 0, -5), *feature.Vector(1, 1, 1)), 0, 0, 0},
	}
	for _, table := range tables {
		c := feature.NewCylinder()
		count, inters, _ := c.IntersectWithRay(&table.r)
		if table.count != 2 && count != table.count {
			t.Errorf("Error Input")
			continue
		}
		if (table.count == 2) && (inters[0].Position != table.ans1 || inters[1].Position != table.ans2) {
			t.Errorf("Error Input, %v, %v", inters[1].Position, table.ans1)
		}
	}
}

//TestCylinder2 tests to see if the NormalAt function works for feature Cylinder
func TestCylinder2(t *testing.T) {
	tables := []struct {
		point  feature.Tuple
		normal feature.Tuple
	}{
		{*feature.Point(1, 0, 0), *feature.Vector(1, 0, 0)},
		{*feature.Point(0, 5, -1), *feature.Vector(0, 0, -1)},
		{*feature.Point(0, -2, 1), *feature.Vector(0, 0, 1)},
		{*feature.Point(-1, 1, 0), *feature.Vector(-1, 0, 0)},
	}
	for _, table := range tables {
		c := feature.NewCylinder()
		normal := c.NormalAt(&table.point)
		if normal != table.normal {
			t.Errorf("Error input")
		}
	}
}
