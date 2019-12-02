package feature

import "math"

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
	deter, _ := to.Transform.Determinant()
	iT := to.Transform.GetInverse(deter)
	newR := ray.Transform(iT)
	circleRaidusSquared := math.Pow(to.CircleRadius, 2)

	//Converting a torus equation with some intermediate variables
	a := 4 * circleRaidusSquared * (math.Pow(newR.Direction.X, 2) + math.Pow(newR.Direction.Y, 2))
	b := 8 * circleRaidusSquared * (newR.Direction.X*newR.Origin.X + newR.Direction.Y*newR.Origin.Y)
	c := 4 * circleRaidusSquared * (math.Pow(newR.Origin.X, 2) + math.Pow(newR.Origin.Y, 2))
	d, _ := newR.Direction.MagnitudeSquared()
	e, _ := newR.Origin.DotProduct(&newR.Direction)
	e = 2 * e
	ftemp, _ := newR.Origin.MagnitudeSquared()
	f := ftemp + circleRaidusSquared - math.Pow(to.TubeRadius, 2)

	//Defining the quartic equation in the following form:
	alpha := d * d
	beta := 2 * d * f
	gamma := 2*d*f + e*e - a
	epsilon := 2*e*f - b
	omega := f*f - c

	//Solving the quartic equation with some intermediate variables
	l := 3*beta - 8*alpha*gamma
	m := -beta + 4*alpha*beta*gamma - 8*alpha*epsilon
	n := 3*beta + 16*(alpha*gamma-alpha*beta*gamma+alpha*beta*epsilon-alpha*omega)
	o := l - 3*n
	p := l*n - 9*m
	q := n - 3*l*m
	delta := p*p - 4*o*q
	if l == 0 && l == m && l == n {
		inter := *NewIntersection(-beta/(4*alpha), *ray, to)
		ans = append(ans, inter)
		return 1, ans, true
	}
	if (l*m*n) != 0 && o == 0 && o == p && o == q {
		inter1 := *NewIntersection((-beta*l+9*m)/(4*alpha*l), *ray, to)
		inter2 := *NewIntersection((-beta*l-3*m)/(4*alpha*l), *ray, to)
		ans = append(ans, inter1, inter2)
		return 2, ans, true
	}
	if m == 0 && m == n && l > 0 {
		inter1 := *NewIntersection((-b+math.Sqrt(l))/(4*alpha), *ray, to)
		inter2 := *NewIntersection((-b*l-3*m)/(4*alpha*l), *ray, to)
		ans = append(ans, inter1, inter2)
		return 2, ans, true
	}
	if (o*p*q) != 0 && delta == 0 && o*p > 0 {
		ans1 := (-beta + sgn(o*p*m)*math.Sqrt(l-p/o) + math.Sqrt(2*p/o)) / (4 * alpha)
		ans2 := (-beta + sgn(o*p*m)*math.Sqrt(l-p/o) - math.Sqrt(2*p/o)) / (4 * alpha)
		ans3 := (-beta + sgn(o*p*m)*math.Sqrt(l-p/o)) / (4 * alpha)
		inter1 := *NewIntersection(ans1, *ray, to)
		inter2 := *NewIntersection(ans2, *ray, to)
		inter3 := *NewIntersection(ans3, *ray, to)
		ans = append(ans, inter1, inter2, inter3)
		return 3, ans, true
	}
	if delta > 0 {
		z1 := o*l + (3*(-p+math.Sqrt(delta)))/2
		z2 := o*l + (3*(-p-math.Sqrt(delta)))/2
		z := l*l - l*(math.Pow(z1, 1/3)+math.Pow(z2, 1/3)) + math.Pow((math.Pow(z1, 1/3)+math.Pow(z2, 1/3)), 2) - 3*o
		pie1 := sgn(m) * math.Sqrt((l+math.Pow(z1, 1/3)+math.Pow(z2, 1/3))/3)
		pie2 := math.Sqrt(2*l - (math.Pow(z1, 1/3) + math.Pow(z2, 1/3)) + 2*math.Sqrt(z))
		pie3 := 4 * alpha
		ans1 := (-beta + pie1 + pie2) / pie3
		ans2 := (-beta + pie1 - pie2) / pie3
		inter1 := *NewIntersection(ans1, *ray, to)
		inter2 := *NewIntersection(ans2, *ray, to)
		ans = append(ans, inter1, inter2)
		return 2, ans, true
	}
	if delta < 0 && l > 0 && n > 0 {
		var ans1, ans2, ans3, ans4 float64
		theta := math.Acos((3*p - 2*o*l) / 2 * math.Pow(o, 3/2))
		y1 := (l - 2*math.Sqrt(o)*math.Cos(theta/3)) / 3
		y2 := (l + math.Sqrt(o)*math.Cos(theta/3) + math.Sqrt(3)*math.Sin(theta/3)) / 3
		y3 := (l + math.Sqrt(o)*math.Cos(theta/3) - math.Sqrt(3)*math.Sin(theta/3)) / 3
		if m != 0 {
			ans1 = (-beta + sgn(m)*math.Sqrt(y1) + math.Sqrt(y2) + math.Sqrt(y3)) / (4 * alpha)
			ans2 = (-beta + sgn(m)*math.Sqrt(y1) - math.Sqrt(y2) - math.Sqrt(y3)) / (4 * alpha)
			ans3 = (-beta - sgn(m)*math.Sqrt(y1) + math.Sqrt(y2) - math.Sqrt(y3)) / (4 * alpha)
			ans4 = (-beta - sgn(m)*math.Sqrt(y1) - math.Sqrt(y2) + math.Sqrt(y3)) / (4 * alpha)
		} else {
			pie1 := math.Sqrt(l + 2*math.Sqrt(n))
			pie2 := math.Sqrt(l - 2*math.Sqrt(n))
			ans1 = (beta + pie1) / (4 * alpha)
			ans2 = (beta - pie1) / (4 * alpha)
			ans3 = (beta + pie2) / (4 * alpha)
			ans4 = (beta - pie2) / (4 * alpha)
		}
		inter1 := *NewIntersection(ans1, *ray, to)
		inter2 := *NewIntersection(ans2, *ray, to)
		inter3 := *NewIntersection(ans3, *ray, to)
		inter4 := *NewIntersection(ans4, *ray, to)
		ans = append(ans, inter1, inter2, inter3, inter4)
		return 4, ans, true
	}
	return 0, ans, false
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
