package test

import (
	"feature"
	"fmt"
	"testing"
)

//TestTriangle1 tests to see if the NormalAt function works for feature Triangle
func TestTriangle1(t *testing.T) {
	tables := []struct {
		point  *feature.Tuple
		normal *feature.Tuple
	}{
		{feature.Point(0, 0.5, 0), feature.Vector(0, 0, -1)},
		{feature.Point(-0.5, 0.75, 0), feature.Vector(0, 0, -1)},
		{feature.Point(0.5, 0.25, 0), feature.Vector(0, 0, -1)},
	}
	for _, table := range tables {
		p1 := feature.Point(0, 1, 0)
		p2 := feature.Point(-1, 0, 0)
		p3 := feature.Point(1, 0, 0)
		tri := feature.NewTriangle(p1, p2, p3)
		normal := tri.NormalAt(table.point)
		if normal != *table.normal {
			fmt.Println(normal)
			fmt.Println(table.normal)
			t.Errorf("Error Input")
		}
	}
}

//TestTriangle2 tests to see if the IntersectWithRay function works for feature Triangle
func TestTriangle2(t *testing.T) {
	tables := []struct {
		ray   *feature.Ray
		count int
		ans   float64
	}{
		{feature.NewRay(*feature.Point(0, -1, -2), *feature.Vector(0, 1, 0)), 0, 0},
		{feature.NewRay(*feature.Point(1, 1, -2), *feature.Vector(0, 0, 1)), 0, 0},
		{feature.NewRay(*feature.Point(-1, 1, -2), *feature.Vector(0, 0, 1)), 0, 0},
		{feature.NewRay(*feature.Point(0, -1, -2), *feature.Vector(0, 0, 1)), 0, 0},
		{feature.NewRay(*feature.Point(0, 0.5, -2), *feature.Vector(0, 0, 1)), 1, 2},
	}
	for _, table := range tables {
		p1 := feature.Point(0, 1, 0)
		p2 := feature.Point(-1, 0, 0)
		p3 := feature.Point(1, 0, 0)
		tri := feature.NewTriangle(p1, p2, p3)
		count, ans, _ := tri.IntersectWithRay(table.ray)
		if count != table.count {
			t.Errorf("Error Input")
		} else if count != 0 && ans[0].Position != table.ans {
			t.Errorf("Hi Error Input")
		}
	}
}
