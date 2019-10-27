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
		typ        string
		ans1, ans2 float64
	}{
		{*feature.NewRay(*feature.Point(1, 0, -5), *feature.Vector(0, 0, 1)), 2, "", 5, 5},
		{*feature.NewRay(*feature.Point(0, 0, -5), *feature.Vector(0, 0, 1)), 2, "", 4, 6},
		{*feature.NewRay(*feature.Point(0.5, 0, -5), *feature.Vector(0.1, 1, 1)), 2, "", 6.80798191702732, 7.088723439378861},
		//Misses a cylinder
		{*feature.NewRay(*feature.Point(1, 0, 0), *feature.Vector(0, 1, 0)), 0, "", 0, 0},
		{*feature.NewRay(*feature.Point(0, 0, 0), *feature.Vector(0, 1, 0)), 0, "", 0, 0},
		{*feature.NewRay(*feature.Point(0, 0, -5), *feature.Vector(1, 1, 1)), 0, "", 0, 0},
		//Truncated cylinder
		{*feature.NewRay(*feature.Point(0, 1.5, 0), *feature.Vector(0.1, 1, 0)), 0, "trunc", 0, 0},
		{*feature.NewRay(*feature.Point(0, 3, -5), *feature.Vector(0, 0, 1)), 0, "trunc", 0, 0},
		{*feature.NewRay(*feature.Point(0, 0, -5), *feature.Vector(0, 0, 1)), 0, "trunc", 0, 0},
		{*feature.NewRay(*feature.Point(0, 2, -5), *feature.Vector(0, 0, 1)), 0, "trunc", 0, 0},
		{*feature.NewRay(*feature.Point(0, 1, -5), *feature.Vector(0, 0, 1)), 0, "trunc", 0, 0},
		{*feature.NewRay(*feature.Point(0, 1.5, -2), *feature.Vector(0, 0, 1)), 2, "trunc", 1, 3},
		//Capped cylinder
		{*feature.NewRay(*feature.Point(0, 3, 0), *feature.Vector(0, -1, 0)), 2, "cap", 0, 0},
		{*feature.NewRay(*feature.Point(0, 3, -2), *feature.Vector(0, -1, 2)), 2, "cap", 0, 0},
		{*feature.NewRay(*feature.Point(0, 4, -2), *feature.Vector(0, -1, 1)), 2, "cap", 0, 0},
		{*feature.NewRay(*feature.Point(0, 0, -2), *feature.Vector(0, 1, 2)), 2, "cap", 0, 0},
		{*feature.NewRay(*feature.Point(0, -1, -2), *feature.Vector(0, 1, 1)), 2, "cap", 1, 3},
	}
	for _, table := range tables {
		c := feature.NewCylinder()
		if table.typ == "trunc" {
			table.r.Direction, _ = table.r.Direction.Normalize()
			c.Max = 2
			c.Min = 1
		}
		if table.typ == "cap" {
			table.r.Direction, _ = table.r.Direction.Normalize()
			c.Max = 2
			c.Min = 1
			c.Closed = true
		}
		count, inters, _ := c.IntersectWithRay(&table.r)
		if (table.count != 2 || table.typ == "cap") && count != table.count {
			t.Errorf("Error Input %v, %v", table.r, count)
			continue
		}
		if (table.count == 2) && (table.typ == "") && (inters[0].Position != table.ans1 || inters[1].Position != table.ans2) {
			t.Errorf("Error Input, %v, %v", inters[1].Position, table.ans1)
		}
	}
}

//TestCylinder2 tests to see if the NormalAt function works for feature Cylinder
func TestCylinder2(t *testing.T) {
	tables := []struct {
		point  feature.Tuple
		typ    string
		normal feature.Tuple
	}{
		{*feature.Point(1, 0, 0), "", *feature.Vector(1, 0, 0)},
		{*feature.Point(0, 5, -1), "", *feature.Vector(0, 0, -1)},
		{*feature.Point(0, -2, 1), "", *feature.Vector(0, 0, 1)},
		{*feature.Point(-1, 1, 0), "", *feature.Vector(-1, 0, 0)},
		//Closed
		{*feature.Point(0, 1, 0), "cap", *feature.Vector(0, -1, 0)},
		{*feature.Point(0.5, 1, 0), "cap", *feature.Vector(0, -1, 0)},
		{*feature.Point(0, 1, 0.5), "cap", *feature.Vector(0, -1, 0)},
		{*feature.Point(0, 2, 0), "cap", *feature.Vector(0, 1, 0)},
		{*feature.Point(0.5, 2, 0), "cap", *feature.Vector(0, 1, 0)},
		{*feature.Point(0, 2, 0.5), "cap", *feature.Vector(0, 1, 0)},
	}
	for _, table := range tables {
		c := feature.NewCylinder()
		if table.typ == "cap" {
			c.Max = 2
			c.Min = 1
			c.Closed = true
		}
		normal := c.NormalAt(&table.point)
		if normal != table.normal {
			t.Errorf("Error input")
		}
	}
}
