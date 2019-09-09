package main

import (
	"fmt"
	"os"
	"strconv"
)

func FloatToString(num float64) string {
	return strconv.FormatFloat(num, 'f', 6, 64)
}

func IntToString(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

func main() {
	// cone := class.NewCanvas(2, 3)
	// //ctwo := class.NewColor(0.9, 1, 0.1)
	// //vtwo := class.NewTuple(2, 3, 4, 0)
	// fmt.Println(cone.PixelAt(0, 1))
	PPMHeader := "P3\n" + IntToString(12) + " " + IntToString(13) + "\n" + "255\n"
	f, err := os.Create("test.ppm")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(PPMHeader)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
