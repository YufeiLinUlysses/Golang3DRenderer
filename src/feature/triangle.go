package feature

import "math"

/*Triangle type contains all necessary components of a sphere
 *Tirangle inherits from Object*/
type Triangle struct {
	Object
	Point1, Point2, Point3 *Tuple
	Edge1, Edge2           Tuple
	Normal                 Tuple
}

/*NewTriangle creates an instance of Type Triangle
 *NewTriangle takes in three tuples
 *NewTriangle returns a triangle with default object*/
func NewTriangle(p1, p2, p3 *Tuple) *Triangle {
	e1, _ := p2.Subtract(p1)
	e2, _ := p3.Subtract(p1)
	cross, _ := e2.CrossProduct(&e1)
	norm, _ := cross.Normalize()
	tri := &Triangle{
		Point1: p1,
		Point2: p2,
		Point3: p3,
		Edge1:  e1,
		Edge2:  e2,
		Normal: norm,
		Object: *NewObject(),
	}
	return tri
}

/*IntersectWithRay calculates the intersections between a triangle and a ray
 *IntersectWithRay can only be called by a triangle
 *IntersectWithRay takes in a ray
 *IntersectWithRay returns a int, a slice of intersection and a bool */
func (tri *Triangle) IntersectWithRay(ray *Ray) (count int, ans []Intersection, intersect bool) {
	deter, _ := tri.Transform.Determinant()
	iT := tri.Transform.GetInverse(deter)
	newR := ray.Transform(iT)
	dirCrossE2, _ := newR.Direction.CrossProduct(&tri.Edge2)
	det, _ := tri.Edge1.DotProduct(&dirCrossE2)
	if math.Abs(det) < 0.00001 {
		return 0, ans, intersect
	}
	f := float64(1) / det
	p1ToOrigin, _ := ray.Origin.Subtract(tri.Point1)
	dot, _ := p1ToOrigin.DotProduct(&dirCrossE2)
	u := f * dot
	if u < 0 || u > 1 {
		return 0, ans, intersect
	}
	oriCrossE1, _ := p1ToOrigin.CrossProduct(&tri.Edge1)
	dot2, _ := ray.Direction.DotProduct(&oriCrossE1)
	v := f * dot2
	if v < 0 || (u+v) > 1 {
		return 0, ans, intersect
	}
	dot3, _ := tri.Edge2.DotProduct(&oriCrossE1)
	t := f * dot3
	inter := NewIntersection(t, *newR, tri)
	ans = append(ans, *inter)
	return 1, ans, true
}

/*NormalAt finds the normal at a certain point
 *NormalAt can only be called by a sphere
 *NormalAt takes in a tuple
 *NormalAt returns a tuple*/
func (tri *Triangle) NormalAt(point *Tuple) Tuple {
	obNormal := tri.Normal
	wNormal := tri.NormalToWorld(&obNormal)
	return *wNormal
}
