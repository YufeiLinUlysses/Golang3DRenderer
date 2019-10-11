package feature

import (
	"sort"
)

//World type
type World struct {
	Lights  []Light
	Objects []interface{}
}

//NewWorld establishes a new world instance, if nothing is given, it returns nothing
func NewWorld(l []Light, o []interface{}) *World {
	w := &World{
		Lights:  l,
		Objects: o,
	}
	return w
}

//DefaultWorld establishes the default world in the book
func DefaultWorld() *World {
	var lights []Light
	var objects []interface{}

	light := NewLight()
	*light = light.PointLight(*Point(-10, 10, -10), *NewColor(1, 1, 1))
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
	var tempCount int
	var ans []Intersection

	for i := range obj {
		switch v := obj[i].(type) {
		case *Sphere:
			tempCount, ans, _ = v.IntersectWithRay(r)
		case *Plane:
			tempCount, ans, _ = v.IntersectWithRay(r)
		}
		count += tempCount
		if tempCount == 1 {
			points = append(points, ans[0])
		} else if tempCount == 2 {
			points = append(points, ans[0], ans[1])
		} else {
			continue
		}
	}
	sort.Slice(points, func(i, j int) bool { return points[i].T < points[j].T })
	return count, points
}

//ShadeHit gives back the color at the intersection in the world
func (w *World) ShadeHit(comp Computations) (colors Color) {
	var mat Material
	switch v := comp.Shape.(type) {
	case *Sphere:
		mat = v.Mat
	case *Plane:
		mat = v.Mat
	}
	for i := range w.Lights {
		light := w.Lights[i]
		inShadow := w.isShadowed(&light, &comp.OverPoint)
		temp := mat.Lighting(light, comp, inShadow)
		colors = colors.Add(&temp)
	}
	return colors
}

//IsShadowed gives whether an object is in shadow or not
func (w *World) isShadowed(l *Light, point *Tuple) bool {
	v, _ := l.Position.Subtract(point)
	distance, _ := v.Magnitude()
	dir, _ := v.Normalize()
	r := NewRay(*point, dir)
	_, inter := w.IntersectWorld(r)
	hit, hitted := Hit(inter)
	if hitted && hit.T < distance {
		return true
	}
	return false
}

//ColorAt returns the color at a
func (w *World) ColorAt(r *Ray) *Color {
	color := NewColor(0, 0, 0)
	_, inters := w.IntersectWorld(r)
	hitPoint, hitted := Hit(inters)
	if hitted == true {
		comp := hitPoint.PrepareComputation(r)
		*color = w.ShadeHit(comp)
	}
	return color
}
