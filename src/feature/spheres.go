package feature

import "math"

/*Sphere type contains all necessary components of a sphere
 *Sphere inherits from Object*/
type Sphere struct {
	Object
	Radius float64
}

/*NewSphere creates an instance of Type Sphere
 *NewSphere returns a sphere with default object*/
func NewSphere() *Sphere {
	sph := &Sphere{
		Object: *NewObject(),
		Radius: 1,
	}
	return sph
}

/*IntersectWithRay calculates the intersections between a sphere and a ray
 *IntersectWithRay can only be called by a sphere
 *IntersectWithRay takes in a ray
 *IntersectWithRay returns a int, a slice of intersection and a bool */
func (sph *Sphere) IntersectWithRay(ray *Ray) (count int, ans []Intersection, intersect bool) {
	deter, _ := sph.Transform.Determinant()
	iT := sph.Transform.GetInverse(deter)
	newR := ray.Transform(iT)
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
		ans = append(ans, *NewIntersection(-b/(2*a), *ray, sph), *NewIntersection(-b/(2*a), *ray, sph))
		return 1, ans, true
	} else {
		ans = append(ans, *NewIntersection((-b-math.Sqrt(discri))/(2*a), *ray, sph), *NewIntersection((-b+math.Sqrt(discri))/(2*a), *ray, sph))
		return 2, ans, true
	}
}

/*NormalAt finds the normal at a certain point
 *NormalAt can only be called by a sphere
 *NormalAt takes in a tuple
 *NormalAt returns a tuple*/
func (sph *Sphere) NormalAt(point *Tuple) Tuple {
	deter, _ := sph.Transform.Determinant()
	inv := sph.Transform.GetInverse(deter)
	obPoint, _ := inv.MultiplyTuple(point)
	obNormal, _ := obPoint.Subtract(Point(0, 0, 0))
	wNormal, _ := inv.Transpose().MultiplyTuple(&obNormal)
	wNormal.W = 0
	ans, _ := wNormal.Normalize()
	return ans
}
