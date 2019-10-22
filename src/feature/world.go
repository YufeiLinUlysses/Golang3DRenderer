package feature

import (
	"math"
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
func (w *World) ColorAt(r *Ray, remaining int) *Color {
	color := NewColor(0, 0, 0)
	_, inters := w.IntersectWorld(r)
	hitPoint, hitted := Hit(inters)
	if hitted == true {
		comp := hitPoint.PrepareComputation(r, inters)
		*color = w.ShadeHit(comp, remaining)
	}
	return color
}

//ReflectedColor returns a color reflected from the object
func (w *World) ReflectedColor(comps Computations, remaining int) *Color {
	var ref float64

	switch v := comps.Shape.(type) {
	case *Sphere:
		ref = v.Mat.Reflectivity
	case Sphere:
		ref = v.Mat.Reflectivity
	case *Plane:
		ref = v.Mat.Reflectivity
	case Plane:
		ref = v.Mat.Reflectivity
	}

	if ref == 0 || remaining == 0 {
		return NewColor(0, 0, 0)
	}
	reflectRay := NewRay(comps.OverPoint, comps.Reflect)
	color := w.ColorAt(reflectRay, remaining-1)
	*color = color.Multiply(ref)
	return color
}

//RefractedColor returns a color refracted from the object
func (w *World) RefractedColor(comps Computations, remaining int) *Color {
	var transp float64
	var color Color
	switch v := comps.Shape.(type) {
	case *Sphere:
		transp = v.Mat.Transparency
	case Sphere:
		transp = v.Mat.Transparency
	case *Plane:
		transp = v.Mat.Transparency
	case Plane:
		transp = v.Mat.Transparency
	}

	nratio := comps.Refract1 / comps.Refract2
	cosi, _ := comps.Eye.DotProduct(&comps.Normal)
	sin2t := math.Pow(nratio, 2) * (1 - math.Pow(cosi, 2))
	if sin2t > 1 {
		return NewColor(0, 0, 0)
	}
	if transp == 0 {
		return NewColor(0, 0, 0)
	}
	if remaining == 0 {
		return NewColor(0, 0, 0)
	}
	cost := math.Sqrt(1 - sin2t)
	calculate := comps.Eye.Multiply(nratio)
	direct := comps.Normal.Multiply(nratio*cosi - cost)
	direct, _ = direct.Subtract(&calculate)
	refractray := NewRay(comps.UnderPoint, direct)
	color = w.ColorAt(refractray, remaining-1).Multiply(transp)
	return &color
}

//ShadeHit gives back the color at the intersection in the world
func (w *World) ShadeHit(comp Computations, remaining int) (colors Color) {
	var mat Material
	var surface, reflected, refracted Color
	switch v := comp.Shape.(type) {
	case *Sphere:
		mat = v.Mat
	case *Plane:
		mat = v.Mat
	}
	for i := range w.Lights {
		light := w.Lights[i]
		inShadow := w.isShadowed(&light, &comp.OverPoint)
		surface = mat.Lighting(light, comp, inShadow)
		reflected = *w.ReflectedColor(comp, remaining)
		refracted = *w.RefractedColor(comp, remaining)
		temp := surface.Add(&reflected)
		temp = temp.Add(&refracted)
		colors = colors.Add(&temp)
	}
	if mat.Reflectivity > 0 && mat.Transparency > 0 {
		reflectance := comp.Schlick()
		temp := reflected.Multiply(reflectance)
		temp = surface.Add(&temp)
		temp1 := refracted.Multiply(1 - reflectance)
		temp = temp.Add(&temp1)
		return temp
	}
	return colors
}
