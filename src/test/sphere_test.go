package test

import (
	"class"
	"testing"
)

//TestSphere1 tests to see if the IntersectWithRay function works for class tuple
func TestSphere1(t *testing.T) {
	tables := []struct {
		r          class.Ray
		s          class.Sphere
		count      int
		ans1, ans2 float64
		intersect  bool
	}{
		{*class.NewRay(0, 0, -5, 0, 0, 1), *class.NewSphere(), 2, 4, 6, true},
		{*class.NewRay(0, 1, -5, 0, 0, 1), *class.NewSphere(), 1, 5, 5, true},
		{*class.NewRay(0, 2, -5, 0, 0, 1), *class.NewSphere(), 0, 0, 0, false},
		{*class.NewRay(0, 0, 0, 0, 0, 1), *class.NewSphere(), 2, -1, 1, true},
		{*class.NewRay(0, 0, 5, 0, 0, 1), *class.NewSphere(), 2, -6, -4, true},
	}
	for _, table := range tables {
		count, ans1, ans2, intersect := table.s.IntersectWithRay(&table.r)
		if count != table.count || ans1 != table.ans1 || ans2 != table.ans2 || intersect != table.intersect {
			t.Errorf("Error Input")
		}
	}
}
