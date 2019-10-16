package test

import (
	"feature"
	"testing"
)

//TestPattern1 tests to see if the StripeAt function works for feature Pattern
func TestPattern1(t *testing.T) {
	tables := []struct {
		point feature.Tuple
		ans   *feature.Color
	}{
		{*feature.Point(0, 0, 0), feature.NewColor(1, 1, 1)},
		{*feature.Point(0, 1, 0), feature.NewColor(1, 1, 1)},
		{*feature.Point(0, 2, 0), feature.NewColor(1, 1, 1)},
		{*feature.Point(0, 0, 1), feature.NewColor(1, 1, 1)},
		{*feature.Point(0, 0, 2), feature.NewColor(1, 1, 1)},
		{*feature.Point(0.9, 0, 0), feature.NewColor(1, 1, 1)},
		{*feature.Point(1, 0, 0), feature.NewColor(0, 0, 0)},
		{*feature.Point(-0.1, 0, 0), feature.NewColor(0, 0, 0)},
		{*feature.Point(-1, 0, 0), feature.NewColor(0, 0, 0)},
		{*feature.Point(-1.1, 0, 0), feature.NewColor(1, 1, 1)},
	}
	for _, table := range tables {
		matrix := feature.NewMatrix(4, 4)
		m, _ := matrix.GetIdentity()
		white := feature.NewColor(1, 1, 1)
		black := feature.NewColor(0, 0, 0)
		p := feature.NewPattern(*white, *black)
		ans := p.StripeAt(table.point, *m)
		if ans.R != table.ans.R || ans.G != table.ans.G || ans.B != table.ans.B {
			t.Errorf("Error Input")
		}
	}
}
