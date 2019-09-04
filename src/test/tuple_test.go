package test

import (
	"class"
	"testing"
)

//Test to see if the GetTuple function works for class tuple
func TestTuple1(t *testing.T) {
	tables := []struct {
		v       class.Vertex
		x, y, z float64
		point   bool
	}{
		{class.Vertex{4.3, -4.2, 3.1, 1.0}, 4.3, -4.2, 3.1, true},
		{class.Vertex{4.5, -42, 310, 0.0}, 4.5, -42, 310, false},
		{class.Vertex{4.5, -42, 310, 0.0}, 4.5, -42, 310, false},
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

func TestTuple2(t *testing.T) {
	tables := []struct {
		v       class.Vertex
		x, y, z float64
		point   bool
	}{
		{class.Vertex{4.3, -4.2, 3.1, 1.0}, 4.3, -4.2, 3.1, true},
		{class.Vertex{4.5, -42, 310, 0.0}, 4.5, -42, 310, false},
		{class.Vertex{4.5, -42, 310, 0.0}, 4.5, -42, 310, false},
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
