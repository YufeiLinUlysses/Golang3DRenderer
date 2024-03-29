package feature

import (
	"sort"
)

/*Group type contains necessary component of a group
 *Group inherits Object*/
type Group struct {
	Object
	Objects []interface{}
}

/*NewGroup gives a default group
 *NewGroup returns an object*/
func NewGroup() *Group {
	matrix := NewMatrix(4, 4)
	matrix, _ = matrix.GetIdentity()
	g := &Group{
		Object: *NewObject(),
	}
	return g
}

/*AddChild adds a shape to group
 *AddChild can only be called by a group
 *AddChild takes in an interface
 *AddChild returns a group instance*/
func (gr *Group) AddChild(shape interface{}) *Group {
	gr.Objects = append(gr.Objects, shape)
	switch v := shape.(type) {
	case *Cube:
		v.Parent = gr
		gr.Objects[len(gr.Objects)-1] = v
	case Cube:
		v.Parent = gr
		gr.Objects[len(gr.Objects)-1] = v
	case *Cylinder:
		v.Parent = gr
		gr.Objects[len(gr.Objects)-1] = v
	case Cylinder:
		v.Parent = gr
		gr.Objects[len(gr.Objects)-1] = v
	case *Cone:
		v.Parent = gr
		gr.Objects[len(gr.Objects)-1] = v
	case Cone:
		v.Parent = gr
		gr.Objects[len(gr.Objects)-1] = v
	case *Sphere:
		v.Parent = gr
		gr.Objects[len(gr.Objects)-1] = v
	case Sphere:
		v.Parent = gr
		gr.Objects[len(gr.Objects)-1] = v
	case *Plane:
		v.Parent = gr
		gr.Objects[len(gr.Objects)-1] = v
	case Plane:
		v.Parent = gr
		gr.Objects[len(gr.Objects)-1] = v
	case *Group:
		v.Parent = gr
		gr.Objects[len(gr.Objects)-1] = v
	case Group:
		v.Parent = gr
		gr.Objects[len(gr.Objects)-1] = v
	case *Triangle:
		v.Parent = gr
		gr.Objects[len(gr.Objects)-1] = v
	case Triangle:
		v.Parent = gr
		gr.Objects[len(gr.Objects)-1] = v
	case *SmoothTriangle:
		v.Parent = gr
		gr.Objects[len(gr.Objects)-1] = v
	case SmoothTriangle:
		v.Parent =gr
		gr.Objects[len(gr.Objects)-1] = v
	case *CSG:
		v.Parent =gr
		gr.Objects[len(gr.Objects)-1] = v
	case CSG:
		v.Parent =gr
		gr.Objects[len(gr.Objects)-1] = v
	case *Torus:
		v.Parent =gr
		gr.Objects[len(gr.Objects)-1] = v
	case Torus:
		v.Parent =gr
		gr.Objects[len(gr.Objects)-1] = v
	}

	return gr
}

/*IntersectWithRay calculates the intersections between the grouped objects and a ray
 *IntersectWithRay can only be called by a group
 *IntersectWithRay takes in a ray
 *IntersectWithRay returns a int, a slice of intersection and a bool */
func (gr *Group) IntersectWithRay(ray *Ray) (count int, ans []Intersection, intersect bool) {
	deter, _ := gr.Transform.Determinant()
	iT := gr.Transform.GetInverse(deter)
	newR := ray.Transform(iT)
	for i := range gr.Objects {
		switch v := gr.Objects[i].(type) {
		case *Cube:
			_, shapeAns, _ := v.IntersectWithRay(newR)
			ans = append(ans, shapeAns...)
		case *Cylinder:
			_, shapeAns, _ := v.IntersectWithRay(newR)
			ans = append(ans, shapeAns...)
		case *Cone:
			_, shapeAns, _ := v.IntersectWithRay(newR)
			ans = append(ans, shapeAns...)
		case *Sphere:
			_, shapeAns, _ := v.IntersectWithRay(newR)
			ans = append(ans, shapeAns...)
		case *Plane:
			_, shapeAns, _ := v.IntersectWithRay(ray)
			ans = append(ans, shapeAns...)
		case *Group:
			_, shapeAns, _ := v.IntersectWithRay(ray)
			ans = append(ans, shapeAns...)
		case *Triangle:
			_, shapeAns, _ := v.IntersectWithRay(ray)
			ans = append(ans, shapeAns...)
		case *SmoothTriangle:
			_, shapeAns, _ := v.IntersectWithRay(ray)
			ans = append(ans, shapeAns...)
		case *CSG:
			_, shapeAns, _ := v.IntersectWithRay(ray)
			ans = append(ans, shapeAns...)
		case *Torus:
			_, shapeAns, _ := v.IntersectWithRay(ray)
			ans = append(ans, shapeAns...)
		}
	}
	sort.Slice(ans, func(i, j int) bool { return ans[i].Position < ans[j].Position })
	return len(ans), ans, len(ans) > 0
}
