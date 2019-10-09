package feature

//Intersection type
type Intersection struct {
	T      float64
	r      Ray
	Object interface{}
}

//NewIntersection establishes a new intersection instance
func NewIntersection(t float64, r Ray, o interface{}) *Intersection {
	i := &Intersection{
		T:      t,
		r:      r,
		Object: o,
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
