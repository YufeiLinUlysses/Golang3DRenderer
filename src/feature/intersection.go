package feature

import (
	"reflect"
)

/*Intersection type contains all necessary component for an inetersection instance
 *Intersection contains a ray and an object*/
type Intersection struct {
	Position float64
	ray      Ray
	Shape    interface{}
	//Triangles only, represents  a location on the surface of a triangle,
	//related to its corners
	U, V float64
}

/*NewIntersection establishes a new intersection instance
 *NewInstersection takes in a float, a ray and an object
 *NewIntersection returns an intersection*/
func NewIntersection(t float64, ray Ray, shape interface{}) *Intersection {
	i := &Intersection{
		Position: t,
		ray:      ray,
		Shape:    shape,
	}
	return i
}

/*Hit generates the hit point
 *Hit could be called anywhere
 *Hit takes in a slice of intersections
 *Hit returns an intersection and a bool*/
func Hit(inters []Intersection) (hitPoint *Intersection, hitted bool) {
	hitted = false
	smallest := float64(999999)
	for i := range inters {
		if inters[i].Position >= 0 && inters[i].Position <= smallest {
			hitted = true
			smallest = inters[i].Position
			hitPoint = &inters[i]
		} else {
			continue
		}
	}
	return hitPoint, hitted
}

/*ShapeInSlice shows whether an object is in an slice or not
 *ShapeInSlice can be called anywhere
 *ShapeInSlice takes in an object and a slice of object
 *ShapeInSlice returns a int and a bool*/
func ShapeInSlice(shape interface{}, mapofshapes []interface{}) (int, bool) {
	var temp1 interface{}
	switch v := shape.(type) {
	case *Cube:
		temp1 = v
	case Cube:
		temp1 = v
	case *Cylinder:
		temp1 = v
	case Cylinder:
		temp1 = v
	case *Cone:
		temp1 = v
	case Cone:
		temp1 = v
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
		case *Cube:
			if reflect.DeepEqual(temp1, temp2) {
				return k, true
			}
			continue
		case Cube:
			if reflect.DeepEqual(temp1, temp2) {
				return k, true
			}
			continue
		case *Cylinder:
			if reflect.DeepEqual(temp1, temp2) {
				return k, true
			}
			continue
		case Cylinder:
			if reflect.DeepEqual(temp1, temp2) {
				return k, true
			}
			continue
		case *Cone:
			if reflect.DeepEqual(temp1, temp2) {
				return k, true
			}
			continue
		case Cone:
			if reflect.DeepEqual(temp1, temp2) {
				return k, true
			}
			continue
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

/*PrepareComputation converts intersection into a computation instance
 *PrepareComputation could only becalled by an intersection instance
 *PrepareComputation takes in a ray and a slice of intersection
 *PrepareComputation returns a computations instance*/
func (intsec *Intersection) PrepareComputation(ray *Ray, inters []Intersection) Computations {
	var comp Computations
	comp.Position = intsec.Position
	comp.Shape = intsec.Shape
	comp.Point = ray.Position(comp.Position)
	comp.Eye = ray.Direction.Multiply(-1)
	switch v := comp.Shape.(type) {
	case *Cube:
		comp.Normal = v.NormalAt(&comp.Point)
	case *Cylinder:
		comp.Normal = v.NormalAt(&comp.Point)
	case *Cone:
		comp.Normal = v.NormalAt(&comp.Point)
	case *Sphere:
		comp.Normal = v.NormalAt(&comp.Point)
	case *Plane:
		comp.Normal = v.NormalAt(&comp.Point)
	case *SmoothTriangle:
		comp.Normal = v.NormalAt(&comp.Point, inters[0])
	}
	if product, _ := comp.Normal.DotProduct(&comp.Eye); product < 0 {
		comp.Inside = true
		comp.Normal = comp.Normal.Multiply(-1)
	} else {
		comp.Inside = false
	}
	multi := comp.Normal.Multiply(0.00001)
	comp.OverPoint = comp.Point.Add(&multi)
	comp.UnderPoint, _ = comp.Point.Subtract(&multi)

	//Get Reflection Index
	comp.Reflect, _ = ray.Direction.Reflect(&comp.Normal)

	//Get Refraction Index
	var container []interface{}
	for i := range inters {
		if reflect.DeepEqual(inters[i], *intsec) {
			if len(container) == 0 {
				comp.Refract1 = 1
			} else {
				switch v := container[len(container)-1].(type) {
				case *Cube:
					comp.Refract1 = v.Mat.Refractivity
				case *Cylinder:
					comp.Refract1 = v.Mat.Refractivity
				case *Cone:
					comp.Refract1 = v.Mat.Refractivity
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
				case *Cube:
					comp.Refract2 = v.Mat.Refractivity
				case *Cylinder:
					comp.Refract2 = v.Mat.Refractivity
				case *Cone:
					comp.Refract2 = v.Mat.Refractivity
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
