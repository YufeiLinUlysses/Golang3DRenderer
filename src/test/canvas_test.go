package test

import (
	"class"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

//TestCanvas1 tests to see if the GetCanvas function works for class Canvas
func TestCanvas1(t *testing.T) {
	tables := []struct {
		w, h       int
		ansW, ansH int
		ansCanv    [][]class.Color
	}{
		{2, 1, 2, 1, [][]class.Color{
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

//TestCanvas2 tests to see if the WritePixel and PixelAt functions work for class Canvas
func TestCanvas2(t *testing.T) {
	tables := []struct {
		w, h   int
		rw, cl int
		inputC class.Color
		ansC   class.Color
	}{
		{2, 1, 1, 0, *class.NewColor(1, 0, 0), *class.NewColor(1, 0, 0)},
	}
	for _, table := range tables {
		c := class.NewCanvas(table.w, table.h)
		newC := c.WritePixel(table.rw, table.cl, &table.inputC)
		ans := newC.PixelAt(table.rw, table.cl)
		if ans != table.ansC {
			t.Errorf("You are wrong")
		}
	}
}

//TestCanvas3 tests to see if the CanvasToPPM work for class Canvas
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
//will CanvasToPPm still work for class Canvas
func TestCanvas4(t *testing.T) {
	canv := class.NewCanvas(10, 2)
	for i := 0; i < canv.Height; i++ {
		for j := 0; j < canv.Width; j++ {
			canv.Canv[i][j] = *class.NewColor(1, 0.8, 0.6)
		}
	}
	canv.CanvasToPPM("hhh")
	files, err:= ioutil.ReadDir("../testFiles/originalFile")
	fmt.Println(files)
	fmt.Println(err)
	dat, err := os.Open("testFiles/originalFile/test2.ppm")
	b, _ := ioutil.ReadAll(dat)
	ans1 := string(b)
	fmt.Println("ans1:\n" + ans1)
	dat1, _ := os.Open("hhh.ppm")
	b1, _ := ioutil.ReadAll(dat1)
	ans2 := string(b1)
	if ans1 != ans2 {
		t.Errorf("You are wrong")
	}
}
