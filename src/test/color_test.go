package test

import (
	"feature"
	"testing"
)

//TestColor1 tests to see if the GetColor  function works for feature Color
func TestColor1(t *testing.T) {
	tables := []struct {
		r, g, b    float64
		rr, gg, bb float64
	}{
		{1, 2, 3, 1, 2, 3},
	}
	for _, table := range tables {
		c := feature.NewColor(table.r, table.g, table.b)
		r, g, b := c.GetColor()
		if r != table.rr || g != table.gg || b != table.bb {
			t.Errorf("Error Input %f,%f,%f", r, g, b)
		}
	}
}

//TestColor2 tests to see if the Add function works for feature Color
func TestColor2(t *testing.T) {
	tables := []struct {
		r, g, b       float64
		rr, gg, bb    float64
		rrr, ggg, bbb float64
	}{
		{1, 2, 3, 4, 5, 6, 5, 7, 9},
	}
	for _, table := range tables {
		cone := feature.NewColor(table.r, table.g, table.b)
		ctwo := feature.NewColor(table.rr, table.gg, table.bb)
		ans := cone.Add(ctwo)
		if ans.R != table.rrr || ans.G != table.ggg || ans.B != table.bbb {
			t.Error("You are wrong")
		}
	}
}

//TestColor3 tests to see if the Subtract function works for feature Color
func TestColor3(t *testing.T) {
	tables := []struct {
		r, g, b       float64
		rr, gg, bb    float64
		rrr, ggg, bbb float64
	}{
		{1, 2, 3, 4, 5, 6, -3, -3, -3},
	}
	for _, table := range tables {
		cone := feature.NewColor(table.r, table.g, table.b)
		ctwo := feature.NewColor(table.rr, table.gg, table.bb)
		ans := cone.Subtract(ctwo)
		if ans.R != table.rrr || ans.G != table.ggg || ans.B != table.bbb {
			t.Error("You are wrong")
		}
	}
}

//TestColor4 tests to see if the Multiply function works for feature Color
func TestColor4(t *testing.T) {
	tables := []struct {
		r, g, b       float64
		scalar        float64
		rrr, ggg, bbb float64
	}{
		{1, 2, 3, 4, 4, 8, 12},
	}
	for _, table := range tables {
		cone := feature.NewColor(table.r, table.g, table.b)
		ans := cone.Multiply(table.scalar)
		if ans.R != table.rrr || ans.G != table.ggg || ans.B != table.bbb {
			t.Error("You are wrong")
		}
	}
}

//TestColor5 tests to see if the ColorMultiply function works for feature Color
func TestColor5(t *testing.T) {
	tables := []struct {
		r, g, b       float64
		rr, gg, bb    float64
		rrr, ggg, bbb float64
	}{
		{1, 2, 3, 4, 5, 6, 4, 10, 18},
	}
	for _, table := range tables {
		cone := feature.NewColor(table.r, table.g, table.b)
		ctwo := feature.NewColor(table.rr, table.gg, table.bb)
		ans := cone.ColorMultiply(ctwo)
		if ans.R != table.rrr || ans.G != table.ggg || ans.B != table.bbb {
			t.Error("You are wrong")
		}
	}
}
