package feature

import "math"

//Plane type
type Plane struct {
	Object
}

//NewPlane sets up an instance for class Plane
func NewPlane() *Plane {
	p := &Plane{
		Object: *NewObject(),
	}
	return p
}

//NormalAt finds the normal at the surface of the plane
func (p *Plane) NormalAt(point *Tuple) Tuple {
	deter, _ := p.Transform.Determinant()
	inv := p.Transform.GetInverse(deter)
	localNormal := Point(0, 1, 0)
	worldNormal, _ := inv.Transpose().MultiplyTuple(localNormal)
	worldNormal.W = 0
	norm, _ := worldNormal.Normalize()
	return norm
}

//IntersectWithRay calculates the intersection between a plane and a ray
func (p *Plane) IntersectWithRay(r *Ray) (count int, ans []Intersection, intersect bool) {
	if math.Abs(r.Direction.Y) < 0.00001 {
		return 0, ans, false
	}
	ans = append(ans, *NewIntersection(-r.Origin.Y/r.Direction.Y, *r, p))
	return 1, ans, true
}
