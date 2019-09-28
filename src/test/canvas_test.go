package test

import (
	"feature"
	"io/ioutil"
	"os"
	"testing"
)

//TestCanvas1 tests to see if the GetCanvas function works for feature Canvas
func TestCanvas1(t *testing.T) {
	tables := []struct {
		w, h       int
		ansW, ansH int
		ansCanv    [][]feature.Color
	}{
		{2, 1, 2, 1, [][]feature.Color{
			{*feature.NewColor(0, 0, 0), *feature.NewColor(0, 0, 0)},
		},
		},
	}
	for _, table := range tables {
		c := feature.NewCanvas(table.w, table.h)
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

//TestCanvas2 tests to see if the WritePixel and PixelAt functions work for feature Canvas
func TestCanvas2(t *testing.T) {
	tables := []struct {
		w, h   int
		rw, cl int
		inputC feature.Color
		ansC   feature.Color
	}{
		{2, 1, 1, 0, *feature.NewColor(1, 0, 0), *feature.NewColor(1, 0, 0)},
	}
	for _, table := range tables {
		c := feature.NewCanvas(table.w, table.h)
		newC := c.WritePixel(table.rw, table.cl, &table.inputC)
		ans := newC.PixelAt(table.rw, table.cl)
		if ans != table.ansC {
			t.Errorf("You are wrong")
		}
	}
}

//TestCanvas3 tests to see if the CanvasToPPM work for feature Canvas
func TestCanvas3(t *testing.T) {
	dat, _ := os.Open("test.ppm")
	b, _ := ioutil.ReadAll(dat)
	ans1 := string(b)
	dat1, _ := os.Open("hahaha.ppm")
	b1, _ := ioutil.ReadAll(dat1)
	ans2 := string(b1)
	if ans1 != ans2 {
		t.Errorf("You are wrong")
	}
}

//TestCanvas4 tests to see if we change everything
//will CanvasToPPm still work for feature Canvas
func TestCanvas4(t *testing.T) {
	canv := feature.NewCanvas(10, 2)
	for i := 0; i < canv.Height; i++ {
		for j := 0; j < canv.Width; j++ {
			canv.Canv[i][j] = *feature.NewColor(1, 0.8, 0.6)
		}
	}
	canv.CanvasToPPM("testFiles/myFile/hhh")
	dat, _ := os.Open("testFiles/originalFile/test2.ppm")
	b, _ := ioutil.ReadAll(dat)
	ans1 := string(b)
	dat1, _ := os.Open("testFiles/myFile/hhh.ppm")
	b1, _ := ioutil.ReadAll(dat1)
	ans2 := string(b1)
	if ans1 != ans2 {
		t.Errorf("You are wrong")
	}
}
