package feature

import "math"

/*Cube type contains all necessary compnent of a cube
 *Cube inherits from Object*/
type Cube struct {
	Object
	Length float64
}

/*NewCube creates an instance of Type Sphere
 *NewCube returns a sphere with default object*/
func NewCube() *Cube {
	cub := &Cube{
		Object: *NewObject(),
		Length: 1,
	}
	return cub
}

/*IntersectWithRay calculates the intersections between a cube and a ray
 *IntersectWithRay can only be called by a cube
 *IntersectWithRay takes in a ray
 *IntersectWithRay returns a int, a slice of intersection and a bool */
func (cub *Cube) IntersectWithRay(ray *Ray) (count int, ans []Intersection, intersect bool) {
	deter, _ := cub.Transform.Determinant()
	iT := cub.Transform.GetInverse(deter)
	newR := ray.Transform(iT)
	xtmin, xtmax := CheckAxis(newR.Origin.X, newR.Direction.X)
	ytmin, ytmax := CheckAxis(newR.Origin.Y, newR.Direction.Y)
	ztmin, ztmax := CheckAxis(newR.Origin.Z, newR.Direction.Z)
	temptmax := math.Min(xtmax, ytmax)
	temptmin := math.Max(xtmin, ytmin)
	tmax := math.Min(ztmax, temptmax)
	tmin := math.Max(ztmin, temptmin)
	if tmin > tmax {
		return 0, ans, false
	}
	ans = append(ans, *NewIntersection(tmin, *ray, cub), *NewIntersection(tmax, *ray, cub))
	return 2, ans, true
}

/*NormalAt finds the normal at a certain point
 *NormalAt can only be called by a cube
 *NormalAt takes in a tuple
 *NormalAt returns a tuple*/
func (cub *Cube) NormalAt(point *Tuple) Tuple {
	var obNormal Tuple
	deter, _ := cub.Transform.Determinant()
	inv := cub.Transform.GetInverse(deter)
	obPoint, _ := inv.MultiplyTuple(point)
	maxc := math.Max(math.Abs(obPoint.X), math.Abs(obPoint.Y))
	maxc = math.Max(maxc, math.Abs(obPoint.Z))
	if maxc == math.Abs(obPoint.X) {
		obNormal = *Vector(obPoint.X, 0, 0)
	} else if maxc == math.Abs(obPoint.Y) {
		obNormal = *Vector(0, obPoint.Y, 0)
	} else {
		obNormal = *Vector(0, 0, obPoint.Z)
	}
	wNormal, _ := inv.Transpose().MultiplyTuple(&obNormal)
	wNormal.W = 0
	ans, _ := wNormal.Normalize()
	return ans
}
