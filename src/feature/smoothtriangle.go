package feature

import "math"

/*SmoothTriangle type
 **/
type SmoothTriangle struct {
	Object
	Point1, Point2, Point3    Tuple
	Edge1, Edge2              Tuple
	Normal1, Normal2, Normal3 Tuple
}

/*NewSmoothTriangle creates an instance of SmoothTriangle*/
func NewSmoothTriangle(p1, p2, p3, n1, n2, n3 Tuple) *SmoothTriangle {
	e1, _ := p2.Subtract(&p1)
	e2, _ := p3.Subtract(&p1)
	st := &SmoothTriangle{
		Object : *NewObject(),
		Point1:  p1,
		Point2:  p2,
		Point3:  p3,
		Edge1:   e1,
		Edge2:   e2,
		Normal1: n1,
		Normal2: n2,
		Normal3: n3,
	}
	return st
}

/*IntersectionWithUandV gives*/
func (stri *SmoothTriangle) IntersectionWithUandV(ray Ray, position, u, v float64) *Intersection {
	i := &Intersection{
		Position: position,
		ray:      ray,
		Shape:    stri,
		U:        u,
		V:        v,
	}
	return i
}

/*IntersectWithRay calculates the intersections between a triangle and a ray
 *IntersectWithRay can only be called by a triangle
 *IntersectWithRay takes in a ray
 *IntersectWithRay returns a int, a slice of intersection and a bool */
func (stri *SmoothTriangle) IntersectWithRay(ray *Ray) (count int, ans []Intersection, intersect bool) {
	deter, _ := stri.Transform.Determinant()
	iT := stri.Transform.GetInverse(deter)
	newR := ray.Transform(iT)
	dirCrossE2, _ := newR.Direction.CrossProduct(&stri.Edge2)
	det, _ := stri.Edge1.DotProduct(&dirCrossE2)
	if math.Abs(det) < 0.00001 {
		return 0, ans, intersect
	}
	f := float64(1) / det
	p1ToOrigin, _ := ray.Origin.Subtract(&stri.Point1)
	dot, _ := p1ToOrigin.DotProduct(&dirCrossE2)
	u := f * dot
	if u < 0 || u > 1 {
		return 0, ans, intersect
	}
	oriCrossE1, _ := p1ToOrigin.CrossProduct(&stri.Edge1)
	dot2, _ := ray.Direction.DotProduct(&oriCrossE1)
	v := f * dot2
	if v < 0 || (u+v) > 1 {
		return 0, ans, intersect
	}
	dot3, _ := stri.Edge2.DotProduct(&oriCrossE1)
	t := f * dot3
	inter := stri.IntersectionWithUandV(*newR, t, u, v)
	ans = append(ans, *inter)
	return 1, ans, true
}

/*NormalAt finds*/
func (stri *SmoothTriangle)NormalAt(point *Tuple, hit Intersection) Tuple{
	temp1 := stri.Normal1.Multiply(1-hit.V-hit.U)
	temp2 := stri.Normal3.Multiply(hit.V)
	temp3 := stri.Normal2.Multiply(hit.U)
	fin := temp2.Add(&temp1)
	obNormal := fin.Add(&temp3)
	wNormal := *stri.NormalToWorld(&obNormal)
	return wNormal
}