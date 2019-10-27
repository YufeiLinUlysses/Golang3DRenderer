package feature

import (
	"math"
)

/*Cone type contains all necessary component of a cone
 *Cone inherits from cylinder*/
type Cone struct {
	Cylinder
}

/*NewCone creates an instance of Type Cone
 *NewCone returns a sphere with default object*/
func NewCone() *Cone {
	con := &Cone{
		Cylinder: *NewCylinder(),
	}
	return con
}

/*IntersectWithRay calculates the intersections between a cone and a ray
 *IntersectWithRay can only be called by a cylinder
 *IntersectWithRay takes in a ray
 *IntersectWithRay returns a int, a slice of intersection and a bool */
func (con *Cone) IntersectWithRay(ray *Ray) (count int, ans []Intersection, intersect bool) {
	yValue := math.Max(con.Max, con.Min)
	deter, _ := con.Transform.Determinant()
	iT := con.Transform.GetInverse(deter)
	newR := ray.Transform(iT)
	a := math.Pow(newR.Direction.X, 2) - math.Pow(newR.Direction.Y, 2) + math.Pow(newR.Direction.Z, 2)
	b := 2*newR.Origin.X*newR.Direction.X - 2*newR.Origin.Y*newR.Direction.Y + 2*newR.Origin.Z*newR.Direction.Z
	c := math.Pow(newR.Origin.X, 2) - math.Pow(newR.Origin.Y, 2) + math.Pow(newR.Origin.Z, 2)
	if math.Abs(a) < 0.00001 && math.Abs(b) < 0.00001 {
		return con.intersectCap(newR, ans, math.Abs(yValue))
	} else if math.Abs(a) < 0.00001 {
		t := -c / (2 * b)
		ans = append(ans, *NewIntersection(t, *ray, con))
	}
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
	if y0 > con.Min && y0 < con.Max {
		ans = append(ans, *NewIntersection(t0, *ray, con))
	}
	y1 := newR.Origin.Y + t1*newR.Direction.Y
	if y1 > con.Min && y1 < con.Max {
		ans = append(ans, *NewIntersection(t1, *ray, con))
	}
	count, ansCap, boolCap := con.intersectCap(newR, ans, math.Abs(yValue))
	return len(ansCap), ansCap, boolCap
}

/*NormalAt finds the normal at a certain point
 *NormalAt can only be called by a cone
 *NormalAt takes in a tuple
 *NormalAt returns a tuple*/
func (con *Cone) NormalAt(point *Tuple) Tuple {
	var obNormal Tuple
	deter, _ := con.Transform.Determinant()
	inv := con.Transform.GetInverse(deter)
	obPoint, _ := inv.MultiplyTuple(point)
	dist := math.Pow(obPoint.X, 2) + math.Pow(obPoint.Z, 2)
	if dist < 1 && obPoint.Y >= con.Max-0.00001 {
		return *Vector(0, 1, 0)
	} else if dist < 1 && obPoint.Y <= con.Min+0.00001 {
		return *Vector(0, -1, 0)
	}
	y := math.Sqrt(math.Pow(obPoint.X, 2) + math.Pow(obPoint.Z, 2))
	if obPoint.Y > 0 {
		y = -y
	}
	obNormal = *Vector(obPoint.X, y, obPoint.Z)
	wNormal, _ := inv.Transpose().MultiplyTuple(&obNormal)
	wNormal.W = 0
	ans, _ := wNormal.Normalize()
	return ans
}
