package feature

//Intersection type
type Intersection struct {
	T     float64
	r     Ray
	Shape interface{}
}

//NewIntersection establishes a new intersection instance
func NewIntersection(t float64, r Ray, o interface{}) *Intersection {
	i := &Intersection{
		T:     t,
		r:     r,
		Shape: o,
	}
	return i
}

//Hit generates the hit point
func Hit(i []Intersection) (hitPoint *Intersection, hitted bool) {
	hitted = false
	smallest := float64(999999)
	for j := range i {
		if i[j].T >= 0 && i[j].T <= smallest {
			hitted = true
			smallest = i[j].T
			hitPoint = &i[j]
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
	}
	if product, _ := comp.Normal.DotProduct(&comp.Eye); product < 0 {
		comp.Inside = true
		comp.Normal = comp.Normal.Multiply(-1)
	} else {
		comp.Inside = false
	}
	multi := comp.Normal.Multiply(0.00001)
	comp.OverPoint = comp.Point.Add(&multi)
	return comp
}
