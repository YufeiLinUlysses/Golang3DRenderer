package class

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

//Canvas type
type Canvas struct {
	Width, Height int
	Canv          [][]Color
}

//WriteErrorFile writes a file for all captured errors in the process
func WriteErrorFile(errorStr string, err error) {
	errorStr += "Have Error: " + err.Error() + "\n"
	errorFile, _ := os.Create("Error.txt")
	errorFile.WriteString(errorStr)
	errorFile.Close()
}

//FloatToString converts float to string
func FloatToString(num float64) string {
	return strconv.FormatFloat(num, 'f', 6, 64)
}

//IntToString converts int to string
func IntToString(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

//ConvertToNum converts pixel number to float between 0 and 255
func ConvertToNum(num float64) int {
	if num >= 255 {
		num = 255
	} else if num <= 0 {
		num = 0
	}
	return int(math.Round(num))
}

//NewCanvas establishes a new Canvas instance
func NewCanvas(w, h int) *Canvas {
	//Initialize the canvas as a 2D slice
	matrix := make([][]Color, h)
	for i := 0; i < h; i++ {
		matrix[i] = make([]Color, w)
	}

	c := &Canvas{
		Width:  w,
		Height: h,
		Canv:   matrix,
	}
	return c
}

//GetCanvas gets the element from Canvas instance
func (c *Canvas) GetCanvas() (w, h int, can [][]Color) {
	return c.Width, c.Height, c.Canv
}

//WritePixel writes the color pixel on a specific location
func (c *Canvas) WritePixel(cl, rw int, col *Color) *Canvas {
	c.Canv[rw][cl] = *col
	return c
}

//PixelAt returns the color pixel at a certain location
func (c *Canvas) PixelAt(cl, rw int) Color {
	return c.Canv[rw][cl]
}

//CanvasToString writes to PPM file
func (c *Canvas) CanvasToString() string {
	var ans string
	PPMHeader := "P3\n" + IntToString(c.Width) + " " + IntToString(c.Height) + "\n" + "255\n"
	ans += PPMHeader

	for i := 0; i < c.Height; i++ {
		var temp string
		for j := 0; j < c.Width; j++ {
			var blue string
			red := IntToString(ConvertToNum(c.Canv[i][j].R*255)) + " "
			green := IntToString(ConvertToNum(c.Canv[i][j].G*255)) + " "
			blue = IntToString(ConvertToNum(c.Canv[i][j].B * 255))
			if j == (c.Width - 1) {
				blue = blue + "\n"
			} else {
				blue = blue + " "
			}
			if len(temp)+len(red) <= 70 {
				temp += red
			} else {
				ans += temp + "\n"
				temp = red
			}
			if len(temp)+len(green) <= 70 {
				temp += green
			} else {
				ans += temp + "\n"
				temp = green
			}
			if len(temp)+len(blue) <= 70 {
				temp += blue
			} else {
				ans += temp + "\n"
				temp = blue
			}
		}
		ans += temp
	}

	return ans
}

//CanvasToPPM saves the canvas string to ppm file
func (c *Canvas) CanvasToPPM(title string) {
	ppm := c.CanvasToString()
	fileTitle := title + ".ppm"
	errorStr := "Writing File:" + fileTitle + time.Now().String() + "\n"
	f, err := os.Create(fileTitle)
	if err != nil {
		fmt.Println(err)
		WriteErrorFile(errorStr, err)
		return
	}
	_, err = f.WriteString(ppm)
	if err != nil {
		fmt.Println(err)
		WriteErrorFile(errorStr, err)
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		WriteErrorFile(errorStr, err)
		return
	}
}
