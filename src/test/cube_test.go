package test

import (
	"feature"
	"testing"
)

//TestCube1 tests to see if the IntersectWithRay function works for feature Color
func TestCube1(t *testing.T) {
	tables := []struct {
		r          feature.Ray
		count      int
		ans1, ans2 float64
	}{
		{*feature.NewRay(*feature.Point(5, 0.5, 0), *feature.Vector(-1, 0, 0)), 2, 4, 6},
		{*feature.NewRay(*feature.Point(-5, 0.5, 0), *feature.Vector(1, 0, 0)), 2, 4, 6},
		{*feature.NewRay(*feature.Point(0.5, 5, 0), *feature.Vector(0, -1, 0)), 2, 4, 6},
		{*feature.NewRay(*feature.Point(0.5, -5, 0), *feature.Vector(0, 1, 0)), 2, 4, 6},
		{*feature.NewRay(*feature.Point(0.5, 0, 5), *feature.Vector(0, 0, -1)), 2, 4, 6},
		{*feature.NewRay(*feature.Point(0.5, 0, -5), *feature.Vector(0, 0, 1)), 2, 4, 6},
		{*feature.NewRay(*feature.Point(0, 0.5, 0), *feature.Vector(0, 0, 1)), 2, -1, 1},
		//Misses a cube
		{*feature.NewRay(*feature.Point(-2, 0, 0), *feature.Vector(0.2673, 0.5345, 0.8010)), 0, 0, 0},
		{*feature.NewRay(*feature.Point(0, -2, 0), *feature.Vector(0.8018, 0.2673, 0.5345)), 0, 0, 0},
		{*feature.NewRay(*feature.Point(0, 0, -2), *feature.Vector(0.5345, 0.8018, 0.2673)), 0, 0, 0},
		{*feature.NewRay(*feature.Point(2, 0, 2), *feature.Vector(0, 0, -1)), 0, 0, 0},
		{*feature.NewRay(*feature.Point(0, 2, 2), *feature.Vector(0, -1, 0)), 0, 0, 0},
		{*feature.NewRay(*feature.Point(2, 2, 0), *feature.Vector(-1, 0, 0)), 0, 0, 0},
	}
	for _, table := range tables {
		c := feature.NewCube()
		count, inters, _ := c.IntersectWithRay(&table.r)
		if table.count != 2 && count != table.count {
			t.Errorf("Error Input")
			continue
		}
		if (table.count == 2) && (inters[0].Position != table.ans1 || inters[1].Position != table.ans2) {
			t.Errorf("Error Input")
		}
	}
}

//TestCube2 tests to see if the NormalAt function works for feature Color
func TestCube2(t *testing.T) {
	tables := []struct {
		point  feature.Tuple
		normal feature.Tuple
	}{
		{*feature.Point(1, 0.5, -0.8), *feature.Vector(1, 0, 0)},
		{*feature.Point(-1, -0.2, 0.9), *feature.Vector(-1, 0, 0)},
		{*feature.Point(-0.4, 1, -0.1), *feature.Vector(0, 1, 0)},
		{*feature.Point(0.3, -1, -0.7), *feature.Vector(0, -1, 0)},
		{*feature.Point(-0.6, 0.3, 1), *feature.Vector(0, 0, 1)},
		{*feature.Point(0.4, 0.4, -1), *feature.Vector(0, 0, -1)},
		{*feature.Point(1, 1, 1), *feature.Vector(1, 0, 0)},
		{*feature.Point(-1, -1, -1), *feature.Vector(-1, 0, 0)},
	}
	for _, table := range tables {
		c := feature.NewCube()
		normal := c.NormalAt(&table.point)
		if normal != table.normal {
			t.Errorf("Error input")
		}
	}
}
