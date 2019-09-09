package class

import(
	"strconv"
)

//Canvas type
type Canvas struct {
	Width, Height int
	Canv          [][]Color
}

//Convert float to string
func FloatToString(num float64) string {
	return strconv.FormatFloat(num, 'f', 6, 64)
}

//Convert int to string
func IntToString(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

//NewCanvas establishes a new Canvas instance
func NewCanvas(w, h int) *Canvas {
	//Initialize the canvas as a 2D slice
	matrix := make([][]Color, w)
	for i := 0; i < w; i++ {
		matrix[i] = make([]Color, h)
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
func (c *Canvas) WritePixel(rw, cl int, col *Color) *Canvas {
	c.Canv[rw][cl] = *col
	return c
}

//PixelAt returns the color pixel at a certain location
func (c *Canvas) PixelAt(rw, cl int) Color {
	return c.Canv[rw][cl]
}

//CanvasToPPM writes to PPM file
func (c *Canvas) CanvasToPPM() string{
	PPMHeader := "P3\n" + IntToString(c.Width) + " " + IntToString(c.Height) + "\n" + "255\n"

}
