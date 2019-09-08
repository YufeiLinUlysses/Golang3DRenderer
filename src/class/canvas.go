package class

//Canvas type
type Canvas struct {
	Width, Height int
	Canv          [][]Color
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
