package test

import (
	"feature"
	"testing"
)

//TestMaterial1 tests to see if the GetMaterial function works for feature Material
func TestMaterial1(t *testing.T) {
	tables := []struct {
		m    feature.Material
		ansc feature.Color
		ansd float64
	}{
		{*feature.NewMaterial(), *feature.NewColor(1, 1, 1), 1},
	}
	for _, table := range tables {
		ansc, ansd := table.m.GetMaterial()
		if ansc != table.ansc || ansd != table.ansd {
			t.Errorf("You are wrong")
		}
	}
}

//TestMaterial2 tests to see if the Lighting function works for feature Material
func TestMaterial2(t *testing.T) {
	tables := []struct {
		m    feature.Material
		s    feature.Sphere
		l    feature.Light
		p    feature.Tuple
		c    feature.Color
		ansc feature.Color
	}{
		{*feature.NewMaterial(), *feature.NewSphere(), *feature.NewLight(), *feature.Point(0, 0, -10), *feature.NewColor(1, 1, 1), *feature.NewColor(1, 1, 1)},
		{*feature.NewMaterial(), *feature.NewSphere(), *feature.NewLight(), *feature.Point(0, 10, -10), *feature.NewColor(1, 1, 1), *feature.NewColor(0.7071067811865475, 0.7071067811865475, 0.7071067811865475)},
		{*feature.NewMaterial(), *feature.NewSphere(), *feature.NewLight(), *feature.Point(0, 0, 10), *feature.NewColor(1, 1, 1), *feature.NewColor(0, 0, 0)},
	}
	for _, table := range tables {
		normal := table.s.NormalAt(feature.Point(0, 0, -1))
		light := table.l.PointLight(table.p, table.c)
		ansc := table.m.Lighting(light, &table.s.Center, &normal)
		if ansc != table.ansc {
			t.Errorf("You are wrong")
		}
	}
}
