package test

import (
	"class"
	"testing"
)

//TestCanvas1 tests to see if the GetCanvas function works for class tuple
func TestCanvas1(t *testing.T) {
	tables := []struct {
		w, h       int
		ansW, ansH int
		ansCanv    [][]class.Color
	}{
		{1, 2, 1, 2, [][]class.Color{
			{*class.NewColor(0, 0, 0), *class.NewColor(0, 0, 0)},
		},
		},
	}
	for _, table := range tables {
		c := class.NewCanvas(table.w, table.h)
		w, h, canv := c.GetCanvas()
		if w != table.ansW || h != table.ansH {
			t.Errorf("Error Input %v,%v,%v", w, h, canv)
		}
		for i, row := range canv {
			for j := range row {
				if canv[i][j] != table.ansCanv[i][j] {
					t.Errorf("Error Input %v,%v,%v", w, h, canv)
				}
			}
		}
	}
}

//TestCanvas2 tests to see if the WritePixel and PixelAt functions work for class tuple
func TestCanvas2(t *testing.T) {
	tables := []struct {
		w, h   int
		rw, cl int
		inputC class.Color
		ansC   class.Color
	}{
		{1, 2, 0, 1, *class.NewColor(1, 0, 0), *class.NewColor(1, 0, 0)},
	}
	for _, table := range tables {
		c := class.NewCanvas(table.w, table.h)
		newC := c.WritePixel(table.rw, table.cl, &table.inputC)
		ans := newC.PixelAt(table.rw,table.cl)
		if ans != table.ansC{
			t.Errorf("You are wrong")
		}
	}
}

//TestCanvas3 tests to see if the WritePixel and PixelAt functions work for class tuple
func TestCanvas3(t *testing.T) {
	tables := []struct {
		w, h   int
		rw, cl int
		inputC class.Color
		ansC   class.Color
	}{
		{1, 2, 0, 1, *class.NewColor(1, 0, 0), *class.NewColor(1, 0, 0)},
	}
	for _, table := range tables {
		c := class.NewCanvas(table.w, table.h)
		newC := c.WritePixel(table.rw, table.cl, &table.inputC)
		ans := newC.PixelAt(table.rw,table.cl)
		if ans != table.ansC{
			t.Errorf("You are wrong")
		}
	}
}
