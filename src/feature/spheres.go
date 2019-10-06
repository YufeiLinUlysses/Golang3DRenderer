package feature

import "math"

//Sphere type, for now just create an unit sphere
type Sphere struct {
	Transform *Matrix
	Material  Material
	Center    Tuple
	Radius    float64
}

//NewSphere creates an instance of Type Sphere
func NewSphere() *Sphere {
	matrix := NewMatrix(4, 4)
	m, _ := matrix.GetIdentity()
	s := &Sphere{
		Transform: m,
		Material:  *NewMaterial(),
		Center:    *Point(0, 0, 0),
		Radius:    1,
	}
	return s
}

//IntersectWithRay calculates the intersection between a sphere and a ray
func (s *Sphere) IntersectWithRay(r *Ray) (count int, ans1, ans2 float64, intersect bool) {
	deter, _ := s.Transform.Determinant()
	iT := s.Transform.GetInverse(deter)
	newR := r.Transform(iT)
	o, _ := newR.Origin.Subtract(NewTuple(0, 0, 0, 1))
	a, _ := newR.Direction.MagnitudeSquared()
	b, _ := newR.Direction.DotProduct(&o)
	b = b * 2
	c, _ := o.MagnitudeSquared()
	c = c - 1
	discri := b*b - 4*a*c
	if discri < 0 {
		return count, ans1, ans2, false
	} else if discri == 0 {
		return 1, (-b / 2), (-b / 2), true
	} else {
		ans1 = (-b - math.Sqrt(discri)) / (2 * a)
		ans2 = (-b + math.Sqrt(discri)) / (2 * a)
		return 2, ans1, ans2, true
	}
}

//NormalAt finds the normal at a certain point
func (s *Sphere) NormalAt(point *Tuple) Tuple {
	ans, _ := point.Subtract(Point(0, 0, 0))
	ans, _ = ans.Normalize()
	return ans
}

//SetTransform sets the Transform variable for feature Sphere
func (s *Sphere) SetTransform(matrix *Matrix) *Sphere {
	s.Transform = matrix
	return s
}
