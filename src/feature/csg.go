package feature

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
func (csg *CSG)IntersectionAllowed(op string, lhit, inl, inr bool) bool{
	if op == "union"{
		return (lhit && !inr) || (!lhit && !inl)
	}
	if op == "intersection"{
		return (lhit && inr)||(!lhit&&inl)
	}
	if op == "difference"{
		return (lhit&&!inr) ||(!lhit&&inl)
	}
	return false
}