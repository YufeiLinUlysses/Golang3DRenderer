package feature

import "sort"

//World type
type World struct {
	Light   []Light
	Objects []interface{}
}

//NewWorld establishes a new world instance, if nothing is given, it returns nothing
func NewWorld(l []Light, o []interface{}) *World {
	w := &World{
		Light:   l,
		Objects: o,
	}
	return w
}

//DefaultWorld establishes the default world in the book
func DefaultWorld() *World {
	var lights []Light
	var objects []interface{}

	light := NewLight()
	lights = append(lights, *light)

	s1 := NewSphere()
	s1.Mat.Col = *NewColor(0.8, 1.0, 0.6)
	s1.Mat.Diffuse = 0.7
	s1.Mat.Specular = 0.2
	objects = append(objects, s1)

	s2 := NewSphere()
	s2.Transform = Scale(0.5, 0.5, 0.5)
	objects = append(objects, s2)

	w := NewWorld(lights, objects)
	return w
}

//IntersectWorld gives the intersection the ray has with the world
func (w *World) IntersectWorld(r *Ray) (count int, points []Intersection) {
	obj := w.Objects
	for i := range obj {
		switch v := obj[i].(type) {
		case *Sphere:
			tempCount, ans, _ := v.IntersectWithRay(r)
			count += tempCount
			if tempCount == 1 {
				points = append(points, ans[0])
			} else if tempCount == 2 {
				points = append(points, ans[0], ans[1])
			} else {
				continue
			}

		}
	}
	sort.Slice(points, func(i, j int) bool { return points[i].T < points[j].T })
	return count, points
}
