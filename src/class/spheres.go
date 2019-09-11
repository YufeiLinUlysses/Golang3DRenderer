package class

import (
	"math"
)

//Canvas type, for now just create an unit sphere
type Sphere struct {
	Center Tuple
	Radius float64
}

func NewSphere() *Sphere {
	s := &Sphere{
		Center: *NewTuple(0, 0, 0, 1),
		Radius: 1,
	}
	return s
}

func (s *Sphere) IntersectWithRay(r *Ray) (count int, ans1, ans2 float64, intersect bool) {
	o, _ := r.Origin.Subtract(NewTuple(0, 0, 0, 1))
	b, _ := r.Direction.DotProduct(&o)
	b = b * 2
	c, _ := o.MagnitudeSquared()
	c = c - 1
	discri := b*b - 4*c
	if discri < 0 {
		return count, ans1, ans2, false
	} else if discri == 0 {
		return 1, (-b / 2), (-b / 2), true
	} else {
		ans1 = (-b - math.Sqrt(discri)) / 2
		ans2 = (-b + math.Sqrt(discri)) / 2
		return 2, ans1, ans2, true
	}
}
