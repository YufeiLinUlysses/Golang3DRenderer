package feature

import (
	"math"
)

//Computations type
type Computations struct {
	T          float64
	Refract1   float64
	Refract2   float64
	Shape      interface{}
	Point      Tuple
	OverPoint  Tuple
	UnderPoint Tuple
	Eye        Tuple
	Normal     Tuple
	Reflect    Tuple
	Inside     bool
}

//Schlick returns the reflectance that is between 0 and 1 and represents what fraction of the light is reflected
func (comp *Computations) Schlick() float64 {
	cos, _ := comp.Eye.DotProduct(&comp.Normal)
	if comp.Refract1 > comp.Refract2 {
		n := comp.Refract1 / comp.Refract2
		sin2t := math.Pow(n, 2) * (1 - math.Pow(cos, 2))
		if sin2t > 1 {
			return 1
		}
		cost := math.Sqrt(1 - sin2t)
		cos = cost
	}
	r0 := math.Pow(((comp.Refract1 - comp.Refract2) / (comp.Refract1 + comp.Refract2)), 2)
	r0 = r0 + (1-r0)*math.Pow((1-cos), 5)
	return r0
}
