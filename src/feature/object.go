package feature

import (
	"math"
)

//Object type
type Object struct {
	Mat       Material
	Center    Tuple
	Transform *Matrix
}

//NewObject sets a new instance for Object class
func NewObject() *Object {
	matrix := NewMatrix(4, 4)
	m, _ := matrix.GetIdentity()
	o := &Object{
		Mat:       *NewMaterial(),
		Center:    *Point(0, 0, 0),
		Transform: m,
	}
	return o
}

//SetTransform sets the transform matrix
func (obj *Object) SetTransform(matrix *Matrix) *Object {
	obj.Transform = matrix
	return obj
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

//ViewTransformation changes the view orientation
func ViewTransformation(from, to, up Tuple) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	subtract, _ := to.Subtract(&from)
	forward, _ := subtract.Normalize()
	upn, _ := up.Normalize()
	left, _ := forward.CrossProduct(&upn)
	trueUp, _ := left.CrossProduct(&forward)
	m = m.Assign(0, 0, left.X)
	m = m.Assign(1, 0, left.Y)
	m = m.Assign(2, 0, left.Z)
	m = m.Assign(0, 1, trueUp.X)
	m = m.Assign(1, 1, trueUp.Y)
	m = m.Assign(2, 1, trueUp.Z)
	m = m.Assign(0, 2, -forward.X)
	m = m.Assign(1, 2, -forward.Y)
	m = m.Assign(2, 2, -forward.Z)
	trans := Translate(-from.X, -from.Y, -from.Z)
	ans, _ := m.Multiply(trans)
	return ans
}
