package feature

import (
	"math"
)

/*Material type contains all necessary components of a material
 *Material contains pattern and color*/
type Material struct {
	Pat          Pattern
	Col          Color
	Ambient      float64
	Diffuse      float64
	Specular     float64
	Shininess    float64
	Reflectivity float64
	Transparency float64
	Refractivity float64
	HasPattern   bool
	PatternType  string
}

/*NewMaterial establishes a default instance for material class
 *NewMaterial returns a material*/
func NewMaterial() *Material {
	m := &Material{
		Col:          *NewColor(1, 1, 1),
		Ambient:      0.1,
		Diffuse:      0.9,
		Specular:     0.9,
		Shininess:    200,
		Reflectivity: float64(0),
		Transparency: 0,
		Refractivity: 1,
		HasPattern:   false,
	}
	return m
}

/*Lighting gets the lighting of the object and decides the color of a pixel using the Phong Model 
 *Lighting can only be called by a material
 *Lighting takes in a light, a computaions, a bool
 *Lighting returns a color*/
func (mat *Material) Lighting(lig Light, comp Computations, isShadow bool) (col Color) {
	var matCol, diffuse, specular, ans Color
	trans := NewMatrix(4, 4)
	black := NewColor(0, 0, 0)

	switch v := comp.Shape.(type) {
	case *Cube:
		trans = v.Transform
	case Cube:
		trans = v.Transform
	case *Cylinder:
		trans = v.Transform
	case Cylinder:
		trans = v.Transform
	case *Cone:
		trans = v.Transform
	case Cone:
		trans = v.Transform
	case *Sphere:
		trans = v.Transform
	case Sphere:
		trans = v.Transform
	case *Plane:
		trans = v.Transform
	case Plane:
		trans = v.Transform
	case *Group:
		trans = v.Transform
	case Group:
		trans = v.Transform
	case *Triangle:
		trans = v.Transform
	case Triangle:
		trans = v.Transform
	case *SmoothTriangle:
		trans = v.Transform
	case SmoothTriangle:
		trans = v.Transform
	case *CSG:
		trans = v.Transform
	case CSG:
		trans = v.Transform
	}

	if mat.HasPattern {
		matCol = *mat.Pat.PatternAt(comp.Point, *trans, mat.PatternType)
	} else {
		matCol = mat.Col
	}

	effectiveCol := matCol.ColorMultiply(&lig.Intensity)
	sub, _ := lig.Position.Subtract(&comp.Point)
	light, _ := sub.Normalize()
	ambient := effectiveCol.Multiply(mat.Ambient)
	if isShadow {
		return ambient
	}
	lightDotNormal, _ := light.DotProduct(&comp.Normal)
	if lightDotNormal < 0 {
		diffuse = *black
		specular = *black
	} else {
		diffuse = effectiveCol.Multiply(mat.Diffuse * lightDotNormal)
		negLight := light.Multiply(-1)
		reflect, _ := negLight.Reflect(&comp.Normal)
		reflectDotEye, _ := reflect.DotProduct(&comp.Eye)
		if reflectDotEye <= 0 {
			specular = *black
		} else {
			factor := math.Pow(reflectDotEye, mat.Shininess)
			specular = lig.Intensity.Multiply(mat.Specular * factor)
		}
	}
	ans = ambient.Add(&diffuse)
	ans = ans.Add(&specular)
	return ans
}
