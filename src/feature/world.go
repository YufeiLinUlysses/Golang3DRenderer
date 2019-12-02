package feature

import (
	"math"
	"sort"
)

/*World type contains all necessary components of a world
 *World contains a slice of light and a slice of objects*/
type World struct {
	Lights  []Light
	Objects []interface{}
}

/*NewWorld establishes a new world instance, if nothing is given, it returns nothing
 *NewWorld takes in a slice of light and a slice of shapes
 *NewWorld returns an empty world*/
func NewWorld(ligs []Light, objs []interface{}) *World {
	w := &World{
		Lights:  ligs,
		Objects: objs,
	}
	return w
}

/*DefaultWorld establishes the default world in the book
 *DefaultWorld returns a default world*/
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

/*IntersectWorld gives the intersection the ray has with the world
 *IntersectWorld can only be called by a world
 *IntersectWorld takes in a ray
 *IntersectWorld returns a int and a slice of intersection*/
func (world *World) IntersectWorld(ray *Ray) (count int, points []Intersection) {
	obj := world.Objects
	var tempCount int
	var ans []Intersection

	for i := range obj {
		switch v := obj[i].(type) {
		case *Cube:
			tempCount, ans, _ = v.IntersectWithRay(ray)
		case *Cylinder:
			tempCount, ans, _ = v.IntersectWithRay(ray)
		case *Cone:
			tempCount, ans, _ = v.IntersectWithRay(ray)
		case *Sphere:
			tempCount, ans, _ = v.IntersectWithRay(ray)
		case *Plane:
			tempCount, ans, _ = v.IntersectWithRay(ray)
		case *Group:
			tempCount, ans, _ = v.IntersectWithRay(ray)
		case *Triangle:
			tempCount, ans, _ = v.IntersectWithRay(ray)
		case *SmoothTriangle:
			tempCount, ans, _ = v.IntersectWithRay(ray)
		case *CSG:
			tempCount, ans, _ = v.IntersectWithRay(ray)
		case *Torus:
			tempCount, ans, _ = v.IntersectWithRay(ray)
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
	sort.Slice(points, func(i, j int) bool { return points[i].Position < points[j].Position })
	return count, points
}

/*isShadowed gives whether an object is in shadow or not
 *isShadowed can only be called by a world
 *isShadowed takes in a light and a tuple
 *isShadowed returns a bool*/
func (world *World) isShadowed(lig *Light, point *Tuple) bool {
	v, _ := lig.Position.Subtract(point)
	distance, _ := v.Magnitude()
	dir, _ := v.Normalize()
	r := NewRay(*point, dir)
	_, inter := world.IntersectWorld(r)
	hit, hitted := Hit(inter)
	if hitted && hit.Position < distance {
		return true
	}
	return false
}

/*ColorAt returns the color at a
 *ColorAt can only be called by a world
 *ColorAt takes in a ray and a int
 *ColorAt returns a color*/
func (world *World) ColorAt(ray *Ray, remaining int) *Color {
	color := NewColor(0, 0, 0)
	_, inters := world.IntersectWorld(ray)
	hitPoint, hitted := Hit(inters)
	if hitted == true {
		comp := hitPoint.PrepareComputation(ray, inters)
		*color = world.ShadeHit(comp, remaining)
	}
	return color
}

/*ReflectedColor returns a color reflected from the object
 *ReflectedColor can only be called by a world
 *ReflectedColor takes in a computations and a int
 *ReflectedColor returns a color*/
func (world *World) ReflectedColor(comps Computations, remaining int) *Color {
	var ref float64

	switch v := comps.Shape.(type) {
	case *Cube:
		ref = v.Mat.Reflectivity
	case Cube:
		ref = v.Mat.Reflectivity
	case *Cylinder:
		ref = v.Mat.Reflectivity
	case Cylinder:
		ref = v.Mat.Reflectivity
	case *Cone:
		ref = v.Mat.Reflectivity
	case Cone:
		ref = v.Mat.Reflectivity
	case *Sphere:
		ref = v.Mat.Reflectivity
	case Sphere:
		ref = v.Mat.Reflectivity
	case *Plane:
		ref = v.Mat.Reflectivity
	case Plane:
		ref = v.Mat.Reflectivity
	case *Group:
		ref = v.Mat.Reflectivity
	case Group:
		ref = v.Mat.Reflectivity
	case *Triangle:
		ref = v.Mat.Reflectivity
	case Triangle:
		ref = v.Mat.Reflectivity
	case *SmoothTriangle:
		ref = v.Mat.Reflectivity
	case SmoothTriangle:
		ref = v.Mat.Reflectivity
	case *CSG:
		ref = v.Mat.Reflectivity
	case CSG:
		ref = v.Mat.Reflectivity
	case *Torus:
		ref = v.Mat.Reflectivity
	case Torus:
		ref = v.Mat.Reflectivity
	}

	if ref == 0 || remaining == 0 {
		return NewColor(0, 0, 0)
	}
	reflectRay := NewRay(comps.OverPoint, comps.Reflect)
	color := world.ColorAt(reflectRay, remaining-1)
	*color = color.Multiply(ref)
	return color
}

/*RefractedColor returns a color refracted from the object
 *RefractedColor can only be called by a world
 *RefreactedColor takes in a computations and a int*/
func (world *World) RefractedColor(comps Computations, remaining int) *Color {
	var transp float64
	var color Color
	switch v := comps.Shape.(type) {
	case *Cube:
		transp = v.Mat.Transparency
	case Cube:
		transp = v.Mat.Transparency
	case *Cylinder:
		transp = v.Mat.Transparency
	case Cylinder:
		transp = v.Mat.Transparency
	case *Cone:
		transp = v.Mat.Transparency
	case Cone:
		transp = v.Mat.Transparency
	case *Sphere:
		transp = v.Mat.Transparency
	case Sphere:
		transp = v.Mat.Transparency
	case *Plane:
		transp = v.Mat.Transparency
	case Plane:
		transp = v.Mat.Transparency
	case *Group:
		transp = v.Mat.Transparency
	case Group:
		transp = v.Mat.Transparency
	case *Triangle:
		transp = v.Mat.Transparency
	case Triangle:
		transp = v.Mat.Transparency
	case *SmoothTriangle:
		transp = v.Mat.Transparency
	case SmoothTriangle:
		transp = v.Mat.Transparency
	case *CSG:
		transp = v.Mat.Transparency
	case CSG:
		transp = v.Mat.Transparency
	case *Torus:
		transp = v.Mat.Transparency
	case Torus:
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
	color = world.ColorAt(refractray, remaining-1).Multiply(transp)
	return &color
}

/*ShadeHit gives back the color at the intersection in the world
 *ShadeHit can only be called by a world
 *ShadeHit takes in a computations and a int
 *ShadeHit returns a color*/
func (world *World) ShadeHit(comp Computations, remaining int) (colors Color) {
	var mat Material
	var surface, reflected, refracted Color
	switch v := comp.Shape.(type) {
	case *Cube:
		mat = v.Mat
	case *Cylinder:
		mat = v.Mat
	case *Cone:
		mat = v.Mat
	case *Sphere:
		mat = v.Mat
	case *Plane:
		mat = v.Mat
	case *Group:
		mat = v.Mat
	case *Triangle:
		mat = v.Mat
	case *SmoothTriangle:
		mat = v.Mat
	case *CSG:
		mat = v.Mat
	case *Torus:
		mat = v.Mat
	}
	for i := range world.Lights {
		light := world.Lights[i]
		inShadow := world.isShadowed(&light, &comp.OverPoint)
		surface = mat.Lighting(light, comp, inShadow)
		reflected = *world.ReflectedColor(comp, remaining)
		refracted = *world.RefractedColor(comp, remaining)
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
