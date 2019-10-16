package feature

import "math"

//Pattern Type
type Pattern struct {
	Transform *Matrix
	A, B Color
}

//NewPattern establishes a new pattern instance
func NewPattern(a, b Color) *Pattern {
	matrix := NewMatrix(4, 4)
	m, _ := matrix.GetIdentity()
	p := &Pattern{
		A: a,
		B: b,
		Transform: m,
	}
	return p
}

//StripeAt returns the color at a certain point
func (p *Pattern) StripeAt(point Tuple, m Matrix) *Color {
	deterM, _ := m.Determinant()
	invM := m.GetInverse(deterM)
	objectPoint, _ := invM.MultiplyTuple(&point)
	deterP, _ := p.Transform.Determinant()
	invP := p.Transform.GetInverse(deterP)
	patternPoint, _ := invP.MultiplyTuple(objectPoint)
	tempX := patternPoint.X
	if math.Mod(math.Floor(tempX), 2) != 0 {
		return &p.B
	}
	return &p.A
}
