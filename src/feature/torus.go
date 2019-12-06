package feature

import (
	"math"
)

/*Torus type contains all necessary components of a torus
 *Torus inherits from Object*/
type Torus struct {
	Object
	CircleRadius, TubeRadius float64
}

/*NewTorus creates an instance of Type Torus
 *NewTorus returns a torus with default object*/
func NewTorus(circle, tube float64) *Torus {
	to := &Torus{
		Object:       *NewObject(),
		CircleRadius: circle,
		TubeRadius:   tube,
	}
	return to
}

/*IntersectWithRay calculates the intersections between a torus and a ray
 *IntersectWithRay can only be called by a torus
 *IntersectWithRay takes in a ray
 *IntersectWithRay returns a int, a slice of intersection and a bool */
func (to *Torus) IntersectWithRay(ray *Ray) (count int, ans []Intersection, intersect bool) {
	coefs := make([]float64, 5)
	inter := *NewIntersection(0, *ray, to)

	deter, _ := to.Transform.Determinant()
	iT := to.Transform.GetInverse(deter)
	newR := ray.Transform(iT)

	circleRadiusSquared := math.Pow(to.CircleRadius, 2)
	tubeRadiusSquared := math.Pow(to.TubeRadius, 2)
	ox := newR.Origin.X
	oy := newR.Origin.Y
	oz := newR.Origin.Z

	dx := newR.Direction.X
	dy := newR.Direction.Y
	dz := newR.Direction.Z

	dirMag, _ := newR.Direction.MagnitudeSquared()
	orMag, _ := newR.Origin.MagnitudeSquared()
	e := orMag - circleRadiusSquared - tubeRadiusSquared
	f := ox*dx + oy*dy + oz*dz
	fourCRS := 4 * circleRadiusSquared

	//Defining the quartic equation in the following form:
	coefs[0] = e*e - fourCRS*(tubeRadiusSquared-oy*oy)
	coefs[1] = 4.0*f*e + 2.0*fourCRS*oy*dy
	coefs[2] = 2.0*dirMag*e + 4.0*f*f + fourCRS*dy*dy
	coefs[3] = 4 * dirMag * f
	coefs[4] = dirMag * dirMag

	root := SolveQuartic(coefs)
	if root.Count>0{
		for i := 0; i < root.Count; i++ {
			inter.Position = root.Ans[i]
			ans = append(ans, inter)
		}
	}
	return root.Count, ans, root.HasRoot
}

/*NormalAt finds the normal at a certain point
 *NormalAt can only be called by a torus
 *NormalAt takes in a tuple
 *NormalAt returns a tuple*/
func (to *Torus) NormalAt(point *Tuple) Tuple {
	obPoint := to.WorldToObject(point)
	ratio := 1 - to.CircleRadius/math.Sqrt(obPoint.X*obPoint.X+obPoint.Y*obPoint.Y)
	nx := ratio * obPoint.X
	ny := ratio * obPoint.Y
	nz := obPoint.Z
	obNormal := Vector(nx, ny, nz)
	wNormal := to.NormalToWorld(obNormal)
	return *wNormal
}

/*sgn is a helper function for finding roots of a quartic equation*/
func sgn(num float64) float64 {
	if num > 0 {
		return float64(1)
	} else if num == 0 {
		return float64(0)
	}
	return float64(-1)
}
