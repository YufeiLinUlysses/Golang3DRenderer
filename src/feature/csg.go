package feature

import "sort"

/*CSG type contains all necessary compnent of a csg
 *CSG inherits from Object and contains two shapes in interface{}*/
type CSG struct {
	Object
	Operation   string
	Left, Right interface{}
}

/*NewCSG creates an instance of Type CSG
 *NewCSG returns a csg with default object*/
func NewCSG(oper string, left, right interface{}) *CSG {
	csg := &CSG{
		Object:    *NewObject(),
		Operation: oper,
		Left:      left,
		Right:     right,
	}
	switch v := left.(type) {
	case *Cube:
		v.Parent = csg
		v.ParentType = "CSG"
	case Cube:
		v.Parent = csg
		v.ParentType = "CSG"
	case *Cylinder:
		v.Parent = csg
		v.ParentType = "CSG"
	case Cylinder:
		v.Parent = csg
		v.ParentType = "CSG"
	case *Cone:
		v.Parent = csg
		v.ParentType = "CSG"
	case Cone:
		v.Parent = csg
		v.ParentType = "CSG"
	case *Sphere:
		v.Parent = csg
		v.ParentType = "CSG"
	case Sphere:
		v.Parent = csg
		v.ParentType = "CSG"
	case *Plane:
		v.Parent = csg
		v.ParentType = "CSG"
	case Plane:
		v.Parent = csg
		v.ParentType = "CSG"
	case *Group:
		v.Parent = csg
		v.ParentType = "CSG"
	case Group:
		v.Parent = csg
		v.ParentType = "CSG"
	case *Triangle:
		v.Parent = csg
		v.ParentType = "CSG"
	case Triangle:
		v.Parent = csg
		v.ParentType = "CSG"
	case *SmoothTriangle:
		v.Parent = csg
		v.ParentType = "CSG"
	case SmoothTriangle:
		v.Parent = csg
		v.ParentType = "CSG"
	case *CSG:
		v.Parent = csg
		v.ParentType = "CSG"
	case CSG:
		v.Parent = csg
		v.ParentType = "CSG"
	case Torus:
		v.Parent = csg
		v.ParentType = "CSG"
	case *Torus:
		v.Parent = csg
		v.ParentType = "CSG"
	}
	switch v := right.(type) {
	case *Cube:
		v.Parent = csg
		v.ParentType = "CSG"
	case Cube:
		v.Parent = csg
		v.ParentType = "CSG"
	case *Cylinder:
		v.Parent = csg
		v.ParentType = "CSG"
	case Cylinder:
		v.Parent = csg
		v.ParentType = "CSG"
	case *Cone:
		v.Parent = csg
		v.ParentType = "CSG"
	case Cone:
		v.Parent = csg
		v.ParentType = "CSG"
	case *Sphere:
		v.Parent = csg
		v.ParentType = "CSG"
	case Sphere:
		v.Parent = csg
		v.ParentType = "CSG"
	case *Plane:
		v.Parent = csg
		v.ParentType = "CSG"
	case Plane:
		v.Parent = csg
		v.ParentType = "CSG"
	case *Group:
		v.Parent = csg
		v.ParentType = "CSG"
	case Group:
		v.Parent = csg
		v.ParentType = "CSG"
	case *Triangle:
		v.Parent = csg
		v.ParentType = "CSG"
	case Triangle:
		v.Parent = csg
		v.ParentType = "CSG"
	case *SmoothTriangle:
		v.Parent = csg
		v.ParentType = "CSG"
	case SmoothTriangle:
		v.Parent = csg
		v.ParentType = "CSG"
	case *CSG:
		v.Parent = csg
		v.ParentType = "CSG"
	case CSG:
		v.Parent = csg
		v.ParentType = "CSG"
	case Torus:
		v.Parent = csg
		v.ParentType = "CSG"
	case *Torus:
		v.Parent = csg
		v.ParentType = "CSG"
	}
	return csg
}

/*IntersectionAllowed finds whether the intersection is allowed at an intersection
 *IntersectionAllowed can only be called by a csg
 *IntersectionAllowed takes in a string, three bool
 *IntersectionAllowed returns a bool*/
func (csg *CSG) IntersectionAllowed(op string, lhit, inl, inr bool) bool {
	if op == "union" {
		return (lhit && !inr) || (!lhit && !inl)
	}
	if op == "intersection" {
		return (lhit && inr) || (!lhit && inl)
	}
	if op == "difference" {
		return (lhit && !inr) || (!lhit && inl)
	}
	return false
}

/*FilterIntersection finds the correct intersection of the two shapes in a csg
 *FilterIntersection can only be called by a csg
 *FilterIntersection takes in a []Intersection
 *FilterIntersection returns a []Intersection*/
func (csg *CSG) FilterIntersection(inters []Intersection) []Intersection {
	inl := false
	inr := false
	result := make([]Intersection, 0)
	for i := range inters {
		lhit := includes(csg.Left, inters[i].Shape)
		if csg.IntersectionAllowed(csg.Operation, lhit, inl, inr) {
			result = append(result, inters[i])
		}
		if lhit {
			inl = !inl
		} else {
			inr = !inr
		}
	}
	return result
}

/*IntersectWithRay calculates the intersections between a csg and a ray
 *IntersectWithRay can only be called by a csg
 *IntersectWithRay takes in a ray
 *IntersectWithRay returns a int, a slice of intersection and a bool*/
func (csg *CSG) IntersectWithRay(ray *Ray) (count int, ans []Intersection, intersect bool) {
	result := make([]Intersection, 0)
	deter, _ := csg.Transform.Determinant()
	iT := csg.Transform.GetInverse(deter)
	newR := ray.Transform(iT)
	leftInter := getIntersections(csg.Left, newR)
	rightInter := getIntersections(csg.Right, newR)
	result = append(result, leftInter...)
	result = append(result, rightInter...)
	sort.Slice(result, func(i, j int) bool { return result[i].Position < result[j].Position })
	return len(csg.FilterIntersection(result)), csg.FilterIntersection(result), len(csg.FilterIntersection(result)) > 0
}

/*getIntersections gets the intersection from the shape and the ray
 *getIntersection takes in a interface{} and a ray
 *getIntersection returns a []Intersection*/
func getIntersections(shape interface{}, ray *Ray) []Intersection {
	result := make([]Intersection, 0)
	switch v := shape.(type) {
	case *Cube:
		_, result, _ = v.IntersectWithRay(ray)
	case Cube:
		_, result, _ = v.IntersectWithRay(ray)
	case *Cylinder:
		_, result, _ = v.IntersectWithRay(ray)
	case Cylinder:
		_, result, _ = v.IntersectWithRay(ray)
	case *Cone:
		_, result, _ = v.IntersectWithRay(ray)
	case Cone:
		_, result, _ = v.IntersectWithRay(ray)
	case *Sphere:
		_, result, _ = v.IntersectWithRay(ray)
	case Sphere:
		_, result, _ = v.IntersectWithRay(ray)
	case *Plane:
		_, result, _ = v.IntersectWithRay(ray)
	case Plane:
		_, result, _ = v.IntersectWithRay(ray)
	case *Group:
		_, result, _ = v.IntersectWithRay(ray)
	case Group:
		_, result, _ = v.IntersectWithRay(ray)
	case *Triangle:
		_, result, _ = v.IntersectWithRay(ray)
	case Triangle:
		_, result, _ = v.IntersectWithRay(ray)
	case *SmoothTriangle:
		_, result, _ = v.IntersectWithRay(ray)
	case SmoothTriangle:
		_, result, _ = v.IntersectWithRay(ray)
	case *CSG:
		_, result, _ = v.IntersectWithRay(ray)
	case CSG:
		_, result, _ = v.IntersectWithRay(ray)
	case *Torus:
		_, result, _ = v.IntersectWithRay(ray)
	case Torus:
		_, result, _ = v.IntersectWithRay(ray)
	}
	return result
}

/*includes finds whether shapeA includes shapeB
 *includes takes in two interface{}
 *includes returns a bool*/
func includes(shapeA, shapeB interface{}) bool {
	switch v := shapeA.(type) {
	case *CSG:
		return (includes(v.Left, shapeB)) || (includes(v.Right, shapeB))
	case *Group:
		for i := range v.Objects {
			if includes(v.Objects[i], shapeB) {
				return true
			}
		}
	default:
		return v == shapeB
	}
	return false
}
