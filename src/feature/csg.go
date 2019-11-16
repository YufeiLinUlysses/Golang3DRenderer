package feature

import "sort"

/*CSG type*/
type CSG struct {
	Object
	Operation   string
	Left, Right interface{}
}

/*NewCSG creates a */
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
	}
	return csg
}

/*IntersectionAllowed does*/
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

/*FilterIntersection creates*/
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

/*IntersectWithRay finds*/
func (csg *CSG) IntersectWithRay(ray *Ray) []Intersection {
	result := make([]Intersection, 0)
	deter, _ := csg.Transform.Determinant()
	iT := csg.Transform.GetInverse(deter)
	newR := ray.Transform(iT)
	leftInter := getIntersections(csg.Left,newR)
	rightInter := getIntersections(csg.Right,newR)
	result = append(result,leftInter...)
	result = append(result,rightInter...)
	sort.Slice(result, func(i, j int) bool { return result[i].Position < result[j].Position })
	return csg.FilterIntersection(result)
}

/*GetIntersections gets*/
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
	}
	return result
}

/*includes finds*/
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
