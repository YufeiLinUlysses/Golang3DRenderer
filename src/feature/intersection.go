package feature

import (
	"reflect"
)

//Intersection type
type Intersection struct {
	T     float64
	ray   Ray
	Shape interface{}
}

//NewIntersection establishes a new intersection instance
func NewIntersection(t float64, r Ray, s interface{}) *Intersection {
	i := &Intersection{
		T:     t,
		ray:   r,
		Shape: s,
	}
	return i
}

//Hit generates the hit point
func Hit(inters []Intersection) (hitPoint *Intersection, hitted bool) {
	hitted = false
	smallest := float64(999999)
	for i := range inters {
		if inters[i].T >= 0 && inters[i].T <= smallest {
			hitted = true
			smallest = inters[i].T
			hitPoint = &inters[i]
		} else {
			continue
		}
	}
	return hitPoint, hitted
}

//PrepareComputation returns a computations instance
func (intsec *Intersection) PrepareComputation(r *Ray, inters []Intersection) Computations {
	var comp Computations
	comp.T = intsec.T
	comp.Shape = intsec.Shape
	comp.Point = r.Position(comp.T)
	comp.Eye = r.Direction.Multiply(-1)
	switch v := comp.Shape.(type) {
	case *Sphere:
		comp.Normal = v.NormalAt(&comp.Point)
	case *Plane:
		comp.Normal = v.NormalAt(&comp.Point)
	}
	if product, _ := comp.Normal.DotProduct(&comp.Eye); product < 0 {
		comp.Inside = true
		comp.Normal = comp.Normal.Multiply(-1)
	} else {
		comp.Inside = false
	}
	multi := comp.Normal.Multiply(0.00001)
	comp.OverPoint = comp.Point.Add(&multi)
	comp.Reflect, _ = r.Direction.Reflect(&comp.Normal)

	//Get Refraction Index
	var container []interface{}
	for i := range inters {
		if reflect.DeepEqual(inters[i], *intsec) {
			if len(container) == 0 {
				comp.Refract1 = 1
			} else {
				switch v := container[len(container)-1].(type) {
				case *Sphere:
					comp.Refract1 = v.Mat.Refractivity
				case *Plane:
					comp.Refract1 = v.Mat.Refractivity
				}
			}
		}

		if len(container) != 0 {
			temp := inters[i].Shape
			if remove, in := ShapeInSlice(temp, container); in {
				copy(container[remove:], container[remove+1:])
				container[len(container)-1] = ""
				container = container[:len(container)-1]
			} else {
				container = append(container, temp)
			}
		} else {
			container = append(container, inters[i].Shape)
		}

		if reflect.DeepEqual(inters[i], *intsec) {
			if len(container) == 0 {
				comp.Refract2 = 1
			} else {
				switch v := container[len(container)-1].(type) {
				case *Sphere:
					comp.Refract2 = v.Mat.Refractivity
				case *Plane:
					comp.Refract2 = v.Mat.Refractivity
				}
			}
		}
	}

	return comp
}

//ShapeInSlice shows whether an object is in an slice or not
func ShapeInSlice(shape interface{}, mapofshapes []interface{}) (int, bool) {
	var temp1 interface{}
	switch v := shape.(type) {
	case *Sphere:
		temp1 = v
	case Sphere:
		temp1 = v
	case *Plane:
		temp1 = v
	case Plane:
		temp1 = v
	}
	for k := range mapofshapes {
		switch temp2 := mapofshapes[k].(type) {
		case *Sphere:
			if reflect.DeepEqual(temp1, temp2) {
				return k, true
			}
			continue
		case Sphere:
			if reflect.DeepEqual(temp1, temp2) {
				return k, true
			}
			continue
		case *Plane:
			if reflect.DeepEqual(temp1, temp2) {
				return k, true
			}
			continue
		case Plane:
			if reflect.DeepEqual(temp1, temp2) {
				return k, true
			}
			continue
		}
	}
	return 0, false
}
