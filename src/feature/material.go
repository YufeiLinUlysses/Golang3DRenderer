package feature

import (
	"math"
)

//Material type
type Material struct {
	Pat         Pattern
	Col         Color
	Ambient     float64
	Diffuse     float64
	Specular    float64
	Shininess   float64
	HasPattern  bool
	PatternType string
}

//NewMaterial establishes a new instance for material feature
func NewMaterial() *Material {
	m := &Material{
		Col:        *NewColor(1, 1, 1),
		Ambient:    0.1,
		Diffuse:    0.9,
		Specular:   0.9,
		Shininess:  200,
		HasPattern: false,
	}
	return m
}

//GetMaterial gets the information of a material instance
func (m *Material) GetMaterial() (col Color, amb, dif, spe, shi float64) {
	return m.Col, m.Ambient, m.Diffuse, m.Specular, m.Shininess
}

//Lighting gets the lighting of the object and decides the color of a pixel
func (m *Material) Lighting(l Light, comp Computations, isShadow bool) (col Color) {
	var matCol, diffuse, specular, ans Color
	trans := NewMatrix(4, 4)
	black := NewColor(0, 0, 0)

	switch v := comp.Shape.(type) {
	case *Sphere:
		trans = v.Transform
	case Sphere:
		trans = v.Transform
	case *Plane:
		trans = v.Transform
	case Plane:
		trans = v.Transform
	}

	if m.HasPattern {
		matCol = *m.Pat.PatternAt(comp.Point, *trans, m.PatternType)
	} else {
		matCol = m.Col
	}

	effectiveCol := matCol.ColorMultiply(&l.Intensity)
	sub, _ := l.Position.Subtract(&comp.Point)
	light, _ := sub.Normalize()
	ambient := effectiveCol.Multiply(m.Ambient)
	if isShadow {
		return ambient
	}
	lightDotNormal, _ := light.DotProduct(&comp.Normal)
	if lightDotNormal < 0 {
		diffuse = *black
		specular = *black
	} else {
		diffuse = effectiveCol.Multiply(m.Diffuse * lightDotNormal)
		negLight := light.Multiply(-1)
		reflect, _ := negLight.Reflect(&comp.Normal)
		reflectDotEye, _ := reflect.DotProduct(&comp.Eye)
		if reflectDotEye <= 0 {
			specular = *black
		} else {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = l.Intensity.Multiply(m.Specular * factor)
		}
	}
	ans = ambient.Add(&diffuse)
	ans = ans.Add(&specular)
	return ans
}
