package feature

import "math"

//Material type
type Material struct {
	Col       Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

//NewMaterial establishes a new instance for material feature
func NewMaterial() *Material {
	m := &Material{
		Col:       *NewColor(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200,
	}
	return m
}

//GetMaterial gets the information of a material instance
func (m *Material) GetMaterial() (col Color, amb, dif, spe, shi float64) {
	return m.Col, m.Ambient, m.Diffuse, m.Specular, m.Shininess
}

//Lighting gets the lighting of the object and decides the color of a pixel
func (m *Material) Lighting(l Light, comp Computations) (col Color) {
	var diffuse, specular, ans Color
	black := NewColor(0, 0, 0)
	effectiveCol := m.Col.ColorMultiply(&l.Intensity)
	sub, _ := l.Position.Subtract(&comp.Point)
	light, _ := sub.Normalize()
	ambient := effectiveCol.Multiply(m.Ambient)
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
