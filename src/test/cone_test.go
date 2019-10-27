package test

import (
	"feature"
	"math"
	"testing"
)

//TestCone1 tests to see if the IntersectWithRay function works for feature Cone
func TestCone1(t *testing.T) {
	tables := []struct {
		rO, rD     feature.Tuple
		count      int
		typ        string
		ans1, ans2 float64
	}{
		{*feature.Point(0, 0, -5), *feature.Vector(0, 0, 1), 2, "", 5, 5},
		{*feature.Point(0, 0, -5), *feature.Vector(1, 1, 1), 2, "", 8.660254037844386, 8.660254037844386},
		{*feature.Point(1, 1, -5), *feature.Vector(-0.5, -1, 1), 2, "", 4.550055679356349, 49.449944320643645},
		{*feature.Point(0, 0, -1), *feature.Vector(0, 1, 1), 1, "", 0.35355339059327373, 0},
		//Capped cone
		{*feature.Point(0, 0, -5), *feature.Vector(0, 1, 0), 0, "cap", 0, 0},
		{*feature.Point(0, 0, -0.25), *feature.Vector(0, 1, 1), 2, "cap", 0, 0},
		{*feature.Point(0, 0, -0.25), *feature.Vector(0, 1, 0), 4, "cap", 0, 0},
	}
	for _, table := range tables {
		c := feature.NewCone()
		if table.typ == "cap" {
			c.Max = 0.5
			c.Min = -0.5
			c.Closed = true
		}
		dir, _ := table.rD.Normalize()
		r := feature.NewRay(table.rO, dir)
		count, inters, _ := c.IntersectWithRay(r)
		if table.typ != "cap" {
			if (table.count == 1) && (inters[0].Position != table.ans1) {
				t.Errorf("Error Input %v, %v", r, count)
				continue
			}
			if (table.count == 2) && (inters[0].Position != table.ans1 || inters[1].Position != table.ans2) {
				t.Errorf("Error Input, %v, %v", inters[0].Position, table.ans1)
			}
		}
		if count != table.count {
			t.Errorf("Error input")
		}
	}
}

//TestCone2 tests to see if the NormalAt function works for feature Cone
func TestCone2(t *testing.T) {
	tables := []struct {
		point  feature.Tuple
		normal feature.Tuple
	}{
		{*feature.Point(0, 0, 0), *feature.Vector(0, 0, 0)},
		{*feature.Point(1, 1, 1), *feature.Vector(1, -math.Sqrt(2), 1)},
		{*feature.Point(-1, -1, 0), *feature.Vector(-1, 1, 0)},
	}
	for _, table := range tables {
		c := feature.NewCone()
		normal := c.NormalAt(&table.point)
		ansNorm, _ := table.normal.Normalize()
		if normal != ansNorm {
			t.Errorf("Error input %v,%v", normal, table.normal)
		}
	}
}
