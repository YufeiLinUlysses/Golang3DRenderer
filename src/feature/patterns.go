package feature

import (
	"math"
)

/*Pattern Type contains all necessary component of a pattern
 *Pattern contains matrix and color*/
type Pattern struct {
	Transform *Matrix
	ColorA, ColorB      Color
}

/*NewPattern establishes a new pattern instance
 *NewPattern takes in two color
 *NewPattern returns a Pattern*/
func NewPattern(cola, colb Color) *Pattern {
	matrix := NewMatrix(4, 4)
	matrix, _ = matrix.GetIdentity()
	pat := &Pattern{
		ColorA:         cola,
		ColorB:         colb,
		Transform: matrix,
	}
	return pat
}

/*PatternAt returns the color at a certain point
 *PatternAt can only be called by a pattern
 *PatternAt takes in a tuple, a matrix and a string
 *PatternAt returns a color*/
func (pat *Pattern) PatternAt(point Tuple, m Matrix, typ string) *Color {
	var col Color
	deterM, _ := m.Determinant()
	invM := m.GetInverse(deterM)
	objectPoint, _ := invM.MultiplyTuple(&point)
	deterP, _ := pat.Transform.Determinant()
	invP := pat.Transform.GetInverse(deterP)
	patternPoint, _ := invP.MultiplyTuple(objectPoint)
	switch typ{
	case "stripe":
		col = *pat.stripePattern(*patternPoint)
	case "gradient":
		col = *pat.gradientPattern(*patternPoint)
	case "ring":
		col = *pat.ringPattern(*patternPoint)
	case "checker":
		col = *pat.checkerPattern(*patternPoint)
	default:
		col = *NewColor(patternPoint.X, patternPoint.Y, patternPoint.Z)
	}
	return &col
}

/*stripPattern creates a stripe pattern for material
 *stripPattern can only be called by a pattern
 *stripPattern takes in a tuple
 *stripPattern returns a color*/
func (pat *Pattern) stripePattern(point Tuple) *Color {
	if math.Mod(math.Floor(point.X), 2) != 0 {
		return &pat.ColorB
	}
	return &pat.ColorA
}

/*gradientPattern creates a gradient pattern for material
 *gradientPattern can only be called by a pattern
 *gradientPattern takes in a tuple
 *gradientPattern returns a color*/
func (pat *Pattern) gradientPattern(point Tuple) *Color {
	distance := pat.ColorB.Subtract(&pat.ColorA)
	fraction := point.X - math.Floor(point.X)
	ans := distance.Multiply(fraction)
	ans = pat.ColorA.Add(&ans)
	return &ans
}

/*ringPattern creates a ring pattern for material
 *ringPattern can only be called by a pattern
 *ringPattern takes in a tuple
 *ringPattern returns a color*/
func (pat *Pattern) ringPattern(point Tuple) *Color {
	sum := point.X*point.X + point.Z*point.Z
	if math.Mod(math.Floor(math.Sqrt(sum)), 2) == 0 {
		return &pat.ColorA
	}
	return &pat.ColorB
}

/*checkerPattern creates a 3D checker pattern for material
 *checkerPattern can only be called by a pattern
 *checkerPattern takes in a tuple
 *checkerPattern returns a color*/
func (pat *Pattern) checkerPattern(point Tuple) *Color {
	sum := math.Floor(point.X) + math.Floor(point.Y) + math.Floor(point.Z)
	if math.Mod(sum, 2) == 0 {
		return &pat.ColorA
	}
	return &pat.ColorB
}
