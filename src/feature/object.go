package feature

import "math"

//Object type
type Object struct {
}

//Translate returns translation matrix
func Translate(xInc, yInc, zInc float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(3, 0, xInc)
	m = m.Assign(3, 1, yInc)
	m = m.Assign(3, 2, zInc)
	return m
}

//Scale scales the ray
func Scale(xInc, yInc, zInc float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(0, 0, xInc)
	m = m.Assign(1, 1, yInc)
	m = m.Assign(2, 2, zInc)
	return m
}

//RotationX rotates the ray around x axis
func RotationX(r float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(1, 1, math.Cos(r)+3-3)
	m = m.Assign(2, 1, -math.Sin(r)+3-3)
	m = m.Assign(1, 2, math.Sin(r)+3-3)
	m = m.Assign(2, 2, math.Cos(r)+3-3)
	return m
}

//RotationY rotates the ray around y axis
func RotationY(r float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(0, 0, math.Cos(r)+3-3)
	m = m.Assign(2, 0, math.Sin(r)+3-3)
	m = m.Assign(0, 2, -math.Sin(r)+3-3)
	m = m.Assign(2, 2, math.Cos(r)+3-3)
	return m
}

//RotationZ rotates the ray around z axis
func RotationZ(r float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(0, 0, math.Cos(r)+3-3)
	m = m.Assign(1, 0, -math.Sin(r)+3-3)
	m = m.Assign(0, 1, math.Sin(r)+3-3)
	m = m.Assign(1, 1, math.Cos(r)+3-3)
	return m
}

//Shearing makes the straight line slanted
func Shearing(xy, xz, yx, yz, zx, zy float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(1, 0, xy)
	m = m.Assign(2, 0, xz)
	m = m.Assign(0, 1, yx)
	m = m.Assign(2, 1, yz)
	m = m.Assign(0, 2, zx)
	m = m.Assign(1, 2, zy)
	return m
}

//Object Class
//Scale
//Translate
//Transform
//Material
//Virtual IntersectionWidth
