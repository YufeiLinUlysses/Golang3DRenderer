package test

import (
	"feature"
	"testing"
)

//TestPattern1 tests to see if the PatternAt function works for feature Pattern
func TestPattern1(t *testing.T) {
	tables := []struct {
		point        feature.Tuple
		transShape   *feature.Matrix
		transPattern *feature.Matrix
		typ          string
		ans          *feature.Color
	}{
		//None
		{*feature.Point(2, 3, 4), feature.Scale(2, 2, 2), feature.Scale(1, 1, 1), "", feature.NewColor(1, 1.5, 2)},
		{*feature.Point(2, 3, 4), feature.Scale(1, 1, 1), feature.Scale(2, 2, 2), "", feature.NewColor(1, 1.5, 2)},
		{*feature.Point(2.5, 3, 3.5), feature.Scale(2, 2, 2), feature.Translate(0.5, 1, 1.5), "", feature.NewColor(0.75, 0.5, 0.25)},
		//Stripe
		{*feature.Point(0, 0, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "stripe", feature.NewColor(1, 1, 1)},
		{*feature.Point(0, 1, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "stripe", feature.NewColor(1, 1, 1)},
		{*feature.Point(0, 2, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "stripe", feature.NewColor(1, 1, 1)},
		{*feature.Point(0, 0, 1), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "stripe", feature.NewColor(1, 1, 1)},
		{*feature.Point(0, 0, 2), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "stripe", feature.NewColor(1, 1, 1)},
		{*feature.Point(0.9, 0, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "stripe", feature.NewColor(1, 1, 1)},
		{*feature.Point(1, 0, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "stripe", feature.NewColor(0, 0, 0)},
		{*feature.Point(-0.1, 0, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "stripe", feature.NewColor(0, 0, 0)},
		{*feature.Point(-1, 0, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "stripe", feature.NewColor(0, 0, 0)},
		{*feature.Point(-1.1, 0, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "stripe", feature.NewColor(1, 1, 1)},
		{*feature.Point(1.5, 0, 0), feature.Scale(2, 2, 2), feature.Scale(1, 1, 1), "stripe", feature.NewColor(1, 1, 1)},
		{*feature.Point(1.5, 0, 0), feature.Scale(1, 1, 1), feature.Scale(2, 2, 2), "stripe", feature.NewColor(1, 1, 1)},
		{*feature.Point(2.5, 0, 0), feature.Scale(2, 2, 2), feature.Translate(0.5, 0, 0), "stripe", feature.NewColor(1, 1, 1)},
		//Gradient
		{*feature.Point(0.25, 0, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "gradient", feature.NewColor(0.75, 0.75, 0.75)},
		{*feature.Point(0.5, 0, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "gradient", feature.NewColor(0.5, 0.5, 0.5)},
		{*feature.Point(0.75, 0, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "gradient", feature.NewColor(0.25, 0.25, 0.25)},
		//Ring
		{*feature.Point(0, 0, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "ring", feature.NewColor(1, 1, 1)},
		{*feature.Point(1, 0, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "ring", feature.NewColor(0, 0, 0)},
		{*feature.Point(0, 0, 1), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "ring", feature.NewColor(0, 0, 0)},
		{*feature.Point(0.708, 0, 0.708), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "ring", feature.NewColor(0, 0, 0)},
		//Checker repeat in x
		{*feature.Point(0, 0, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "checker", feature.NewColor(1, 1, 1)},
		{*feature.Point(0.99, 0, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "checker", feature.NewColor(1, 1, 1)},
		{*feature.Point(1.01, 0, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "checker", feature.NewColor(0, 0, 0)},
		//Checker repeat in y
		{*feature.Point(0, 0, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "checker", feature.NewColor(1, 1, 1)},
		{*feature.Point(0, 0.99, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "checker", feature.NewColor(1, 1, 1)},
		{*feature.Point(0, 1.01, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "checker", feature.NewColor(0, 0, 0)},
		//Checker repeat in z
		{*feature.Point(0, 0, 0), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "checker", feature.NewColor(1, 1, 1)},
		{*feature.Point(0, 0, 0.99), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "checker", feature.NewColor(1, 1, 1)},
		{*feature.Point(0, 0, 1.01), feature.Scale(1, 1, 1), feature.Scale(1, 1, 1), "checker", feature.NewColor(0, 0, 0)},
	}
	for _, table := range tables {
		s := feature.NewSphere()
		s.Transform = table.transShape
		white := feature.NewColor(1, 1, 1)
		black := feature.NewColor(0, 0, 0)
		p := feature.NewPattern(*white, *black)
		p.Transform = table.transPattern
		ans := p.PatternAt(table.point, *s.Transform, table.typ)
		if ans.R != table.ans.R || ans.G != table.ans.G || ans.B != table.ans.B {
			t.Errorf("Error Input %v", table.point)
		}
	}
}
