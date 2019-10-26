package feature

import "math"

/*Cylinder type contains all necessary compnent of a cylinder
 *Cylinder inherits from Object*/
type Cylinder struct {
	Object
	Radius float64
}

/*NewCylinder creates an instance of Type Sphere
 *NewCylinder returns a sphere with default object*/
func NewCylinder() *Cylinder {
	cyl := &Cylinder{
		Object: *NewObject(),
		Radius: 1,
	}
	return cyl
}

/*IntersectWithRay calculates the intersections between a cylinder and a ray
 *IntersectWithRay can only be called by a cylinder
 *IntersectWithRay takes in a ray
 *IntersectWithRay returns a int, a slice of intersection and a bool */
func (cyl *Cylinder) IntersectWithRay(ray *Ray) (count int, ans []Intersection, intersect bool) {
	deter, _ := cyl.Transform.Determinant()
	iT := cyl.Transform.GetInverse(deter)
	newR := ray.Transform(iT)
	a := math.Pow(newR.Direction.X, 2) + math.Pow(newR.Direction.Z, 2)
	if a < 0.00001 {
		return 0, ans, false
	}
	b := 2*newR.Origin.X*newR.Direction.X + 2*newR.Origin.Z*newR.Direction.Z
	c := math.Pow(newR.Origin.X, 2) + math.Pow(newR.Origin.Z, 2) - 1
	disc := math.Pow(b, 2) - 4*a*c
	if disc < 0 {
		return 0, ans, false
	}
	t0 := ((-b - math.Sqrt(disc)) / (2 * a))
	t1 := ((-b + math.Sqrt(disc)) / (2 * a))
	ans = append(ans, *NewIntersection(t0, *ray, cyl), *NewIntersection(t1, *ray, cyl))
	return 2, ans, true
}

/*NormalAt finds the normal at a certain point
 *NormalAt can only be called by a cylinder
 *NormalAt takes in a tuple
 *NormalAt returns a tuple*/
func (cyl *Cylinder) NormalAt(point *Tuple) Tuple {
	var obNormal Tuple
	deter, _ := cyl.Transform.Determinant()
	inv := cyl.Transform.GetInverse(deter)
	obPoint, _ := inv.MultiplyTuple(point)
	obNormal = *Vector(obPoint.X, 0, obPoint.Z)
	wNormal, _ := inv.Transpose().MultiplyTuple(&obNormal)
	wNormal.W = 0
	ans, _ := wNormal.Normalize()
	return ans
}
