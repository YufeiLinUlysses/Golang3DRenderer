package test

import (
	"class"
	"testing"
)

//TestTuple1 tests to see if the GetTuple function works for class tuple
func TestTuple1(t *testing.T) {
	tables := []struct {
		v       class.Tuple
		x, y, z float64
		point   bool
	}{
		{class.Tuple{4.3, -4.2, 3.1, 1.0}, 4.3, -4.2, 3.1, true},
		{class.Tuple{4.5, -42, 310, 0.0}, 4.5, -42, 310, false},
		{class.Tuple{4.5, -42, 310, 0.0}, 4.5, -42, 310, false},
	}
	for _, table := range tables {
		x, y, z, typeOfTuple := table.v.GetTuple()
		if x != table.x || y != table.y || z != table.z {
			t.Errorf("Error Input")
		} else {
			if !typeOfTuple {
				t.Errorf("Not a point")
			} else {
				t.Errorf("It's a point")
			}
		}
	}
}

//TestTuple2 tests to see if the function Point works as the way we want
func TestTuple2(t *testing.T) {
	tables := []struct {
		x, y, z float64
		ans     class.Tuple
	}{
		{4.3, -4.2, 3.1, class.Tuple{4.3, -4.2, 3.1, 1.0}},
	}
	for _, table := range tables {
		point := class.Point(table.x, table.y, table.z)
		if point != table.ans {
			t.Error("You are wrong")
		}
	}
}

//TestTuple3 tests to see if the function Vector works as the way we want
func TestTuple3(t *testing.T) {
	tables := []struct {
		x, y, z float64
		ans     class.Tuple
	}{
		{4.3, -4.2, 3.1, class.Tuple{4.3, -4.2, 3.1, 0}},
	}
	for _, table := range tables {
		point := class.Vector(table.x, table.y, table.z)
		if point != table.ans {
			t.Error("You are wrong")
		}
	}
}

//TestTuple4 tests to see if the function Add works as the way we want
func TestTuple4(t *testing.T) {
	tables := []struct {
		tone class.Tuple
		ttwo class.Tuple
		ans  class.Tuple
	}{
		{class.Tuple{4.3, -4.2, 3.1, 0}, class.Tuple{4.3, -4.2, 3.1, 1}, class.Tuple{8.6, -8.4, 6.2, 1}},
	}
	for _, table := range tables {
		ans := table.tone.Add(table.ttwo)
		if ans != table.ans {
			t.Error("You are wrong")
		}
	}
}
