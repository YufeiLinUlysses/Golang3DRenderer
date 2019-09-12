package main

import (
	"fmt"
	"path/filepath"
)

//"class"

func main() {
	fmt.Println(filepath.Dir("../test/testFiles/originalFile/"))
	// canv := class.NewCanvas(5, 3)
	// cone := class.NewColor(1.5, 0, 0)
	// ctwo := class.NewColor(0, 0.5, 0)
	// cthree := class.NewColor(-0.5, 0, 1)
	// canv.WritePixel(0, 0, cone)
	// canv.WritePixel(2, 1, ctwo)
	// canv.WritePixel(4, 2, cthree)
	// PPM := canv.CanvasToString()
	// canv.CanvasToPPM("hahaha")
	// fmt.Println(PPM)
	// dat, _ := os.Open("test.ppm")
	// b, _ := ioutil.ReadAll(dat)
	// ans1 := string(b)
	// dat1, _ := os.Open("hahaha.ppm")
	// b1, _ := ioutil.ReadAll(dat1)
	// ans2 := string(b1)
	// fmt.Println(ans1 == ans2)
}
