package test

import (
	"class"
	"testing"
)

//TestLight1 tests to see if the PointLight function works for class Light
func TestLight1(t *testing.T) {
	tables := []struct {
		l    class.Light
		p    class.Tuple
		i    class.Color
		ansp class.Tuple
		ansi class.Color
	}{
		{*class.NewLight(), *class.Point(0, 0, 0), *class.NewColor(1, 1, 1), *class.Point(0, 0, 0), *class.NewColor(1, 1, 1)},
	}
	for _, table := range tables {
		ansl := table.l.PointLight(table.p, table.i)
		if ansl.Position != table.ansp || ansl.Intensity != table.ansi{
			t.Errorf("You are wrong")
		}
	}
}
