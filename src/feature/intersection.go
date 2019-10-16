package feature

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
func (i *Intersection) PrepareComputation(r *Ray) Computations {
	var comp Computations
	comp.T = i.T
	comp.Shape = i.Shape
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
	return comp
}
