package test

import (
	"feature"
	"testing"
)

//TestLight1 tests to see if the PointLight function works for feature Light
func TestLight1(t *testing.T) {
	tables := []struct {
		l    feature.Light
		p    feature.Tuple
		i    feature.Color
		ansp feature.Tuple
		ansi feature.Color
	}{
		{*feature.NewLight(), *feature.Point(0, 0, 0), *feature.NewColor(1, 1, 1), *feature.Point(0, 0, 0), *feature.NewColor(1, 1, 1)},
	}
	for _, table := range tables {
		ansl := table.l.PointLight(table.p, table.i)
		if ansl.Position != table.ansp || ansl.Intensity != table.ansi{
			t.Errorf("You are wrong")
		}
	}
}
