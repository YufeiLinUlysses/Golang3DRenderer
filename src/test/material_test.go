package test

import (
	"class"
	"testing"
)

//TestMaterial1 tests to see if the GetMaterial function works for class Material
func TestMaterial1(t *testing.T) {
	tables := []struct {
		m    class.Material
		ansc class.Color
		ansd float64
	}{
		{*class.NewMaterial(), *class.NewColor(1, 1, 1), 1},
	}
	for _, table := range tables {
		ansc, ansd := table.m.GetMaterial()
		if ansc != table.ansc || ansd != table.ansd {
			t.Errorf("You are wrong")
		}
	}
}

//TestMaterial2 tests to see if the Lighting function works for class Material
func TestMaterial2(t *testing.T) {
	tables := []struct {
		m    class.Material
		s    class.Sphere
		l    class.Light
		p    class.Tuple
		c    class.Color
		ansc class.Color
	}{
		{*class.NewMaterial(), *class.NewSphere(), *class.NewLight(), *class.Point(0, 0, -10), *class.NewColor(1, 1, 1), *class.NewColor(1, 1, 1)},
		{*class.NewMaterial(), *class.NewSphere(), *class.NewLight(), *class.Point(0, 10, -10), *class.NewColor(1, 1, 1), *class.NewColor(0.7071067811865475, 0.7071067811865475, 0.7071067811865475)},
		{*class.NewMaterial(), *class.NewSphere(), *class.NewLight(), *class.Point(0, 0, 10), *class.NewColor(1, 1, 1), *class.NewColor(0, 0, 0)},
	}
	for _, table := range tables {
		normal := table.s.NormalAt(class.Point(0, 0, -1))
		light := table.l.PointLight(table.p, table.c)
		ansc := table.m.Lighting(light, &table.s.Center, &normal)
		if ansc != table.ansc {
			t.Errorf("You are wrong")
		}
	}
}
