package feature

import (
	"math"
)

//Camera type
type Camera struct {
	Hsize, Vsize          int
	HalfWidth, HalfHeight float64
	FieldOfView           float64
	PixelSize             float64
	Transform             *Matrix
}

//NewCamera establishes a new camera instance
func NewCamera(h, v int, f float64) *Camera {
	var halfW, halfH, p float64
	matrix := NewMatrix(4, 4)
	m, _ := matrix.GetIdentity()
	halfView := math.Tan(f / float64(2))
	aspect := float64(h) / float64(v)
	if aspect >= 1 {
		halfW = halfView
		halfH = halfView / aspect
	} else {
		halfW = halfView * aspect
		halfH = halfView
	}
	p = (halfW * 2) / float64(h)
	c := &Camera{
		Hsize:       h,
		Vsize:       v,
		HalfWidth:   halfW,
		HalfHeight:  halfH,
		FieldOfView: f,
		PixelSize:   p,
		Transform:   m,
	}
	return c
}

//RayForPixel gives a new ray that starts at the camera and passes throught the pixel on the canvas
func (c *Camera) RayForPixel(x, y float64) *Ray {
	xOffSet := (x + 0.5) * c.PixelSize
	yOffSet := (y + 0.5) * c.PixelSize
	wx := c.HalfWidth - xOffSet
	wy := c.HalfHeight - yOffSet
	deter, _ := c.Transform.Determinant()
	inv := c.Transform.GetInverse(deter)
	pixel, _ := inv.MultiplyTuple(Point(wx, wy, -1))
	ori, _ := inv.MultiplyTuple(Point(0, 0, 0))
	dirPrep, _ := pixel.Subtract(ori)
	dir, _ := dirPrep.Normalize()
	r := NewRay(*ori, dir)
	return r
}

//Render gives an image of the given world
func (c *Camera) Render(w World) *Canvas {
	image := NewCanvas(c.Hsize, c.Vsize)
	for y := 0; y < c.Vsize; y++ {
		for x := 0; x < c.Hsize; x++ {
			color := w.ColorAt(c.RayForPixel(float64(x), float64(y)),4)
			image.WritePixel(x, y, color)
		}
	}
	return image
}
