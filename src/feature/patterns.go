package feature

import (
	"math"
)

//Pattern Type
type Pattern struct {
	Transform *Matrix
	ColorA, ColorB      Color
}

//NewPattern establishes a new pattern instance
func NewPattern(a, b Color) *Pattern {
	matrix := NewMatrix(4, 4)
	m, _ := matrix.GetIdentity()
	p := &Pattern{
		ColorA:         a,
		ColorB:         b,
		Transform: m,
	}
	return p
}

//PatternAt returns the color at a certain point
func (p *Pattern) PatternAt(point Tuple, m Matrix, typ string) *Color {
	deterM, _ := m.Determinant()
	invM := m.GetInverse(deterM)
	objectPoint, _ := invM.MultiplyTuple(&point)
	deterP, _ := p.Transform.Determinant()
	invP := p.Transform.GetInverse(deterP)
	patternPoint, _ := invP.MultiplyTuple(objectPoint)
	if typ == "stripe" {
		return p.stripePattern(*patternPoint)
	} else if typ == "gradient" {
		return p.gradientPattern(*patternPoint)
	} else if typ == "ring" {
		return p.ringPattern(*patternPoint)
	} else if typ == "checker" {
		return p.checkerPattern(*patternPoint)
	}
	return NewColor(patternPoint.X, patternPoint.Y, patternPoint.Z)
}

//stripPattern creates a stripe pattern for material
func (p *Pattern) stripePattern(point Tuple) *Color {
	if math.Mod(math.Floor(point.X), 2) != 0 {
		return &p.ColorB
	}
	return &p.ColorA
}

//gradientPattern creates a gradient pattern for material
func (p *Pattern) gradientPattern(point Tuple) *Color {
	distance := p.ColorB.Subtract(&p.ColorA)
	fraction := point.X - math.Floor(point.X)
	ans := distance.Multiply(fraction)
	ans = p.ColorA.Add(&ans)
	return &ans
}

//ringPattern creates a ring pattern for material
func (p *Pattern) ringPattern(point Tuple) *Color {
	sum := point.X*point.X + point.Z*point.Z
	if math.Mod(math.Floor(math.Sqrt(sum)), 2) == 0 {
		return &p.ColorA
	}
	return &p.ColorB
}

//checkerPattern creates a 3D checker pattern for material
func (p *Pattern) checkerPattern(point Tuple) *Color {
	sum := math.Floor(point.X) + math.Floor(point.Y) + math.Floor(point.Z)
	if math.Mod(sum, 2) == 0 {
		return &p.ColorA
	}
	return &p.ColorB
}
