package feature

import (
	"math"
)

/*Cylinder type contains all necessary compnent of a cylinder
 *Cylinder inherits from Object*/
type Cylinder struct {
	Object
	Max    float64
	Min    float64
	Closed bool
}

/*NewCylinder creates an instance of Type Cylinder
 *NewCylinder returns a sphere with default object*/
func NewCylinder() *Cylinder {
	cyl := &Cylinder{
		Object: *NewObject(),
		Max:    math.Pow10(10),
		Min:    -math.Pow10(10),
		Closed: false,
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
		return cyl.intersectCap(newR, ans, 1)
	}
	b := 2*newR.Origin.X*newR.Direction.X + 2*newR.Origin.Z*newR.Direction.Z
	c := math.Pow(newR.Origin.X, 2) + math.Pow(newR.Origin.Z, 2) - 1
	disc := math.Pow(b, 2) - 4*a*c
	if disc < 0 {
		return 0, ans, false
	}
	t0 := ((-b - math.Sqrt(disc)) / (2 * a))
	t1 := ((-b + math.Sqrt(disc)) / (2 * a))
	if t0 > t1 {
		t0, t1 = t1, t0
	}
	y0 := newR.Origin.Y + t0*newR.Direction.Y
	if y0 > cyl.Min && y0 < cyl.Max {
		ans = append(ans, *NewIntersection(t0, *ray, cyl))
	}
	y1 := newR.Origin.Y + t1*newR.Direction.Y
	if y1 > cyl.Min && y1 < cyl.Max {
		ans = append(ans, *NewIntersection(t1, *ray, cyl))
	}
	count, ansCap, boolCap := cyl.intersectCap(newR, ans, 1)
	return len(ansCap), ansCap, boolCap
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
	dist := math.Pow(obPoint.X, 2) + math.Pow(obPoint.Z, 2)
	if dist < 1 && obPoint.Y >= cyl.Max-0.00001 {
		return *Vector(0, 1, 0)
	} else if dist < 1 && obPoint.Y <= cyl.Min+0.00001 {
		return *Vector(0, -1, 0)
	}
	obNormal = *Vector(obPoint.X, 0, obPoint.Z)
	wNormal, _ := inv.Transpose().MultiplyTuple(&obNormal)
	wNormal.W = 0
	ans, _ := wNormal.Normalize()
	return ans
}

/*intersectCap finds the intersection of the ray and the cap of the cylinder and cone
 *intersectCap can only be called by a cylinder
 *intersectCap takes in a ray, a slice of intersection and a float
 *intersectCap returns a int, a slice of intersection and a bool*/
func (cyl *Cylinder) intersectCap(ray *Ray, inters []Intersection, radi float64) (count int, ans []Intersection, intersect bool) {
	if cyl.Closed == false || math.Abs(ray.Direction.Y) < 0.00001 {
		if len(inters) == 0 {
			return 0, inters, false
		}
		return len(inters), inters, true
	}
	t := (cyl.Min - ray.Origin.Y) / ray.Direction.Y
	if ray.CheckCap(t, radi) {
		inters = append(inters, *NewIntersection(t, *ray, cyl))
	}
	t = (cyl.Max - ray.Origin.Y) / ray.Direction.Y
	if ray.CheckCap(t, radi) {
		inters = append(inters, *NewIntersection(t, *ray, cyl))
	}
	if len(inters) == 0 {
		return 0, inters, false
	}
	return len(inters), inters, true
}
