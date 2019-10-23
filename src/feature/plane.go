package feature

import "math"

/*Plane type contains all necessary component of a plane
 *Plane inherits from object*/
type Plane struct {
	Object
}

/*NewPlane sets up an instance for class Plane
 *NewPlane returns a plane with a default object*/
func NewPlane() *Plane {
	pl := &Plane{
		Object: *NewObject(),
	}
	return pl
}

/*NormalAt finds the normal at the surface of the plane
 *NormalAt can only be called by a pattern
 *NormalAt takes in a tuple
 *NormalAt returns a tuple*/
func (pl *Plane) NormalAt(point *Tuple) Tuple {
	deter, _ := pl.Transform.Determinant()
	inv := pl.Transform.GetInverse(deter)
	localNormal := Vector(0, 1, 0)
	worldNormal, _ := inv.Transpose().MultiplyTuple(localNormal)
	worldNormal.W = 0
	norm, _ := worldNormal.Normalize()
	return norm
}

/*IntersectWithRay calculates the intersection between a plane and a ray
 *IntersectWithRay can only be called by a plane
 *IntersectWithRay takes in a ray
 *IntersectWithRay returns a int, a slice of intersection and a bool*/
func (pl *Plane) IntersectWithRay(ray *Ray) (count int, ans []Intersection, intersect bool) {
	deter, _ := pl.Transform.Determinant()
	inv := pl.Transform.GetInverse(deter)
	ray = ray.Transform(inv)
	if math.Abs(ray.Direction.Y) < 0.00001 {
		return 0, ans, false
	}
	ans = append(ans, *NewIntersection(-ray.Origin.Y/ray.Direction.Y, *ray, pl))
	return 1, ans, true
}
