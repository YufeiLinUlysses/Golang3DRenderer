package feature

import (
	"math"
)

/*Camera type contains necessary component of a camera
 *Camera contains a matrix*/
type Camera struct {
	Hsize, Vsize          int
	HalfWidth, HalfHeight float64
	FieldOfView           float64
	PixelSize             float64
	Transform             *Matrix
}

/*NewCamera establishes a new camera instance
 *NewCamera takes in two int for creating a canvas and one float for establishing the field of view
 *NewCamera returns a new camera instance*/
func NewCamera(horsize, versize int, fieldofview float64) *Camera {
	var halfW, halfH, pixsize float64

	matrix := NewMatrix(4, 4)
	matrix, _ = matrix.GetIdentity()
	halfView := math.Tan(fieldofview / float64(2))
	aspect := float64(horsize) / float64(versize)
	if aspect >= 1 {
		halfW = halfView
		halfH = halfView / aspect
	} else {
		halfW = halfView * aspect
		halfH = halfView
	}
	pixsize = (halfW * 2) / float64(horsize)
	cam := &Camera{
		Hsize:       horsize,
		Vsize:       versize,
		HalfWidth:   halfW,
		HalfHeight:  halfH,
		FieldOfView: fieldofview,
		PixelSize:   pixsize,
		Transform:   matrix,
	}
	return cam
}

/*RayForPixel gives a new ray that starts at the camera and passes throught the pixel on the canvas
 *It could only be called by a camera instance
 *RayForPixel takes in two float representing the object coordinates
 *RayForPixel returns a new ray*/
func (cam *Camera) RayForPixel(x, y float64) *Ray {
	xOffSet := (x + 0.5) * cam.PixelSize
	yOffSet := (y + 0.5) * cam.PixelSize
	wx := cam.HalfWidth - xOffSet
	wy := cam.HalfHeight - yOffSet
	deter, _ := cam.Transform.Determinant()
	inv := cam.Transform.GetInverse(deter)
	pixel, _ := inv.MultiplyTuple(Point(wx, wy, -1))
	ori, _ := inv.MultiplyTuple(Point(0, 0, 0))
	dirPrep, _ := pixel.Subtract(ori)
	dir, _ := dirPrep.Normalize()
	ray := NewRay(*ori, dir)
	return ray
}

/*Render gives an image of the given world
 *It could only be called by a camera instance
 *Rener takes in a world object that contains all the necessary component of the world
 *Render returns a new canvas object that can be output as a ppm file*/
func (cam *Camera) Render(world World) *Canvas {
	image := NewCanvas(cam.Hsize, cam.Vsize)
	for y := 0; y < cam.Vsize; y++ {
		for x := 0; x < cam.Hsize; x++ {
			color := world.ColorAt(cam.RayForPixel(float64(x), float64(y)), 4)
			image.WritePixel(x, y, color)
		}
	}
	return image
}
