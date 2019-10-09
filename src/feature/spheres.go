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
func (s *Sphere) IntersectWithRay(r *Ray) (count int, ans []Intersection, intersect bool) {
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
		return count, ans, false
	} else if discri == 0 {
		ans = append(ans, *NewIntersection((-b / 2 * a), *r, s), *NewIntersection((-b / 2 * a), *r, s))
		return 1, ans, true
	} else {
		ans = append(ans, *NewIntersection((-b-math.Sqrt(discri))/(2*a), *r, s), *NewIntersection((-b+math.Sqrt(discri))/(2*a), *r, s))
		return 2, ans, true
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
