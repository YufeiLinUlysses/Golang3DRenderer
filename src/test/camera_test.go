package test

import (
	"feature"
	"math"
	"testing"
)

//TestCamera1 tests to see if the NewCamera function works for feature Camera
func TestCamera1(t *testing.T) {
	tables := []struct {
		h, v int
		f    float64
		ans  float64
	}{
		{200, 125, math.Pi / float64(2), 0.01},
		{125, 200, math.Pi / float64(2), 0.01},
	}
	for _, table := range tables {
		c := feature.NewCamera(table.h, table.v, table.f)
		ans := c.PixelSize
		if ans != table.ans {
			t.Errorf("Error Input")
		}
	}
}

//TestCamera2 tests to see if the RayForPixel function works for feature Camera
func TestCamera2(t *testing.T) {
	tables := []struct {
		x, y float64
		ans  *feature.Ray
	}{
		{100, 50, feature.NewRay(*feature.Point(0, 0, 0), *feature.Vector(0, 0, -1))},
		{0, 0, feature.NewRay(*feature.Point(0, 0, 0), *feature.Vector(0.6651864261194509, 0.33259321305972545, -0.6685123582500481))},
	}
	for _, table := range tables {
		c := feature.NewCamera(201, 101, math.Pi/2)
		r := c.RayForPixel(table.x, table.y)
		if r.Origin != table.ans.Origin || r.Direction != table.ans.Direction {
			t.Errorf("Error Input %v, %v", table.ans, r)
		}
	}
}

//TestCamera3 tests to see if the RayForPixel function works for feature Camera with a given transform matrix
func TestCamera3(t *testing.T) {
	tables := []struct {
		x, y float64
		ans  *feature.Ray
	}{
		{100, 50, feature.NewRay(*feature.Point(0, 2, -5), *feature.Vector(0.7071067811865475, 0, -0.7071067811865478))},
	}
	for _, table := range tables {
		c := feature.NewCamera(201, 101, math.Pi/2)
		c.Transform, _ = feature.RotationY(math.Pi / 4).Multiply(feature.Translate(0, -2, 5))
		r := c.RayForPixel(table.x, table.y)
		if *r != *table.ans {
			t.Errorf("Error Input %v, %v", table.ans, r)
		}
	}
}

//TestCamera4 tests to see if the Render function works for feature Camera with a given transform matrix
func TestCamera4(t *testing.T) {
	tables := []struct {
		h, v int
		ans  *feature.Color
	}{
		{11, 11, feature.NewColor(0.38066119308103435, 0.47582649135129296, 0.28549589481077575)},
	}
	for _, table := range tables {
		w := feature.DefaultWorld()
		c := feature.NewCamera(table.h, table.v, math.Pi/2)
		from := feature.Point(0, 0, -5)
		to := feature.Point(0, 0, 0)
		up := feature.Vector(0, 1, 0)
		c.Transform = feature.ViewTransformation(*from, *to, *up)
		image := c.Render(*w)
		if image.PixelAt(5, 5) != *table.ans {
			t.Errorf("Error Input %v, %v", table.ans, image.PixelAt(5, 5))
		}
	}
}
