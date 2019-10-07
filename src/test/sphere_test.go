package test

import (
	"feature"
	"math"
	"testing"
)

//TestSphere1 tests to see if the IntersectWithRay function works for feature tuple
func TestSphere1(t *testing.T) {
	tables := []struct {
		r                feature.Ray
		s                feature.Sphere
		command          string
		xInc, yInc, zInc float64
		count            int
		ans1, ans2       float64
		intersect        bool
	}{
		{*feature.NewRay(*feature.Point(0, 0, -5), *feature.Vector(0, 0, 1)), *feature.NewSphere(), "none", 0, 0, 0, 2, 4, 6, true},
		{*feature.NewRay(*feature.Point(0, 1, -5), *feature.Vector( 0, 0, 1)), *feature.NewSphere(), "none", 0, 0, 0, 1, 5, 5, true},
		{*feature.NewRay(*feature.Point(0, 2, -5), *feature.Vector( 0, 0, 1)), *feature.NewSphere(), "none", 0, 0, 0, 0, 0, 0, false},
		{*feature.NewRay(*feature.Point(0, 0, 0), *feature.Vector(0, 0, 1)), *feature.NewSphere(), "none", 0, 0, 0, 2, -1, 1, true},
		{*feature.NewRay(*feature.Point(0, 0, 5), *feature.Vector( 0, 0, 1)), *feature.NewSphere(), "none", 0, 0, 0, 2, -6, -4, true},
		{*feature.NewRay(*feature.Point(0, 0, -5), *feature.Vector( 0, 0, 1)), *feature.NewSphere(), "scale", 2, 2, 2, 2, 3, 7, true},
		{*feature.NewRay(*feature.Point(0, 0, -5), *feature.Vector( 0, 0, 1)), *feature.NewSphere(), "translate", 5, 0, 0, 0, 0, 0, false},
	}
	for _, table := range tables {
		if table.command == "translate" {
			matrix := feature.Translate(table.xInc, table.yInc, table.zInc)
			table.s = *table.s.SetTransform(matrix)
		} else if table.command == "scale" {
			matrix := feature.Scale(table.xInc, table.yInc, table.zInc)
			table.s = *table.s.SetTransform(matrix)
		}
		count, ans1, ans2, intersect := table.s.IntersectWithRay(&table.r)
		if count != table.count || ans1 != table.ans1 || ans2 != table.ans2 || intersect != table.intersect {
			t.Errorf("Error Input")
		}
	}
}

//TestSphere2 tests to see if the NormalAt function works for feature tuple
func TestSphere2(t *testing.T) {
	tables := []struct {
		s   feature.Sphere
		p   *feature.Tuple
		ans *feature.Tuple
	}{
		{*feature.NewSphere(), feature.Point(1, 0, 0), feature.Vector(1, 0, 0)},
		{*feature.NewSphere(), feature.Point(0, 1, 0), feature.Vector(0, 1, 0)},
		{*feature.NewSphere(), feature.Point(0, 0, 1), feature.Vector(0, 0, 1)},
		{*feature.NewSphere(), feature.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3), feature.Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)},
	}
	for _, table := range tables {
		normal := table.s.NormalAt(table.p)
		if normal != *table.ans {
			t.Errorf("Error Input")
		}
	}
}

//TestSphere3 tests to see if the Transform function works for feature tuple
func TestSphere3(t *testing.T) {
	tables := []struct {
		s   feature.Sphere
		ans []float64
	}{
		{*feature.NewSphere(), []float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}},
	}
	for _, table := range tables {
		count := 0
		ans := table.s.Transform
		for i, row := range ans.Matrix {
			for j := range row {
				if ans.GetValueAt(i, j) != table.ans[count] {
					t.Errorf("wrong %v, %v, %v, %v", ans.GetValueAt(i, j), table.ans[count], i, j)
					break
				}
				count++
			}
		}
	}
}
