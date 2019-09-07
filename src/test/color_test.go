package test

import (
	"class"
	"testing"
)

//TestColor1 tests to see if the GetColor  function works for class tuple
func TestColor1(t *testing.T) {
	tables := []struct {
		r, g, b    float64
		rr, gg, bb float64
	}{
		{1, 2, 3, 1, 2, 3},
	}
	for _, table := range tables {
		c := class.NewColor(table.r, table.g, table.b)
		r, g, b := c.GetColor()
		if r != table.rr || g != table.gg || b != table.bb {
			t.Errorf("Error Input %f,%f,%f", r, g, b)
		}
	}
}

//TestColor2 tests to see if the Add  function works for class tuple
func TestColor2(t *testing.T) {
	tables := []struct {
		r, g, b       float64
		rr, gg, bb    float64
		rrr, ggg, bbb float64
	}{
		{1, 2, 3, 4, 5, 6, 5, 7, 9},
	}
	for _, table := range tables {
		cone := class.NewColor(table.r, table.g, table.b)
		ctwo := class.NewColor(table.rr, table.gg, table.bb)
		ans := cone.Add(ctwo)
		if ans.R != table.rrr || ans.G != table.ggg || ans.B != table.bbb {
			t.Error("You are wrong")
		}
	}
}
