package feature

import "math"

//Sphere type, for now just create an unit sphere
type Sphere struct {
	Object
	Radius float64
}

//NewSphere creates an instance of Type Sphere
func NewSphere() *Sphere {
	s := &Sphere{
		Object: *NewObject(),
		Radius: 1,
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
