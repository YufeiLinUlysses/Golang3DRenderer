package feature

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

/*Canvas type contains necessary component of a canvas
 *Canvas contains a 2D slice of color*/
type Canvas struct {
	Width, Height int
	Canv          [][]Color
}

/*WriteErrorFile writes a file for all captured errors in the process
 *WriteErrorFile takes in a string and an error to form a error text file
 *WriteErrorFile outputs a text file to the computer*/
func WriteErrorFile(errorStr string, err error) {
	errorStr += "Have Error: " + err.Error() + "\n"
	errorFile, _ := os.Create("Error.txt")
	errorFile.WriteString(errorStr)
	errorFile.Close()
}

/*FloatToString converts float to string
 *FloatToString takes in a float
 *FloatToString returns the string conversion of the input number*/
func FloatToString(num float64) string {
	return strconv.FormatFloat(num, 'f', 6, 64)
}

/*IntToString converts int to string
 *IntToString takes in a int
 *IntToString returns the string conversion of the input number*/
func IntToString(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

/*ConvertToNum converts pixel number to float between 0 and 255
 *ConvertToNum takes in a float representing the color
 *ConvertToNum returns a int representing the color that machines can read*/
func ConvertToNum(num float64) int {
	if num >= 255 {
		num = 255
	} else if num <= 0 {
		num = 0
	}
	return int(math.Round(num))
}

/*NewCanvas establishes a new Canvas instance
 *NewCanvas takes in two int representing the width and height of the canvas
 *NewCanvas returns a new canvas*/
func NewCanvas(width, height int) *Canvas {
	//Initialize the canvas as a 2D slice
	matrix := make([][]Color, height)
	for i := 0; i < height; i++ {
		matrix[i] = make([]Color, width)
	}

	canv := &Canvas{
		Width:  width,
		Height: height,
		Canv:   matrix,
	}
	return canv
}

/*GetCanvas gets the element from Canvas instance
 *GetCanvas can only be called by a canvas instance
 *GetCanvas returns the width, height and the canvas itself*/
func (canv *Canvas) GetCanvas() (w, h int, can [][]Color) {
	return canv.Width, canv.Height, canv.Canv
}

/*WritePixel writes the color pixel on a specific location
 *WritePixel can only be called by a canvas instance,
 *WritePixel takes in the two int representing the column and rows of the pixel point and a color
 *WritePixel returns a new canvas*/
func (canv *Canvas) WritePixel(cl, rw int, col *Color) *Canvas {
	canv.Canv[rw][cl] = *col
	return canv
}

/*PixelAt returns the color pixel at a certain location
 *PixelAt could only be called by a canvas instance
 *PixelAt takes in two int representing the column and rows of the pixel point
 *PixelAt returns a color instance*/
func (canv *Canvas) PixelAt(cl, rw int) Color {
	return canv.Canv[rw][cl]
}

/*CanvasToString converts canvas to string for writing to PPM file
 *CanvasToString could only be called by a canvas instance
 *CanvasToString returns a string*/
func (canv *Canvas) CanvasToString() string {
	var ans string
	PPMHeader := "P3\n" + IntToString(canv.Width) + " " + IntToString(canv.Height) + "\n" + "255\n"
	ans += PPMHeader

	for i := 0; i < canv.Height; i++ {
		var temp string
		for j := 0; j < canv.Width; j++ {
			var blue string
			red := IntToString(ConvertToNum(canv.Canv[i][j].R*255)) + " "
			green := IntToString(ConvertToNum(canv.Canv[i][j].G*255)) + " "
			blue = IntToString(ConvertToNum(canv.Canv[i][j].B * 255))
			if j == (canv.Width - 1) {
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

/*CanvasToPPM saves the canvas string to ppm file
 *CanvasToPPM could only be called by a canvas instance
 *CanvasToPPM takes in a string representing the title of the output file
 *CanvasToPPM outputs the file as a PPM file*/
func (canv *Canvas) CanvasToPPM(title string) {
	ppm := canv.CanvasToString()
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
