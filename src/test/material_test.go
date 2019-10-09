package test

import (
	"feature"
	"math"
	"testing"
)

//TestMaterial1 tests to see if the GetMaterial function works for feature Material
func TestMaterial1(t *testing.T) {
	tables := []struct {
		m                       feature.Material
		ansc                    feature.Color
		ansa, ansd, anss, anssh float64
	}{
		{*feature.NewMaterial(), *feature.NewColor(1, 1, 1), 0.1, 0.9, 0.9, 200},
	}
	for _, table := range tables {
		ansc, ansa, ansd, anss, anssh := table.m.GetMaterial()
		if ansc != table.ansc || ansa != table.ansa || ansd != table.ansd || anss != table.anss || anssh != table.anssh {
			t.Errorf("You are wrong")
		}
	}
}

//TestMaterial2 tests to see if the Lighting function works for feature Material
func TestMaterial2(t *testing.T) {
	tables := []struct {
		eye, normal, position *feature.Tuple
		color                 *feature.Color
		ansc                  feature.Color
	}{
		{feature.Vector(0, 0, -1), feature.Vector(0, 0, -1), feature.Point(0, 10, -10), feature.NewColor(1, 1, 1), *feature.NewColor(0.7363961030678927, 0.7363961030678927, 0.7363961030678927)},
		{feature.Vector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2), feature.Vector(0, 0, -1), feature.Point(0, 10, -10), feature.NewColor(1, 1, 1), *feature.NewColor(1.6363961030678928, 1.6363961030678928, 1.6363961030678928)},
		{feature.Vector(0, 0, -1), feature.Vector(0, 0, -1), feature.Point(0, 0, 10), feature.NewColor(1, 1, 1), *feature.NewColor(0.1, 0.1, 0.1)},
		{feature.Vector(0, 0, -1), feature.Vector(0, 0, -1), feature.Point(0, 0, -10), feature.NewColor(1, 1, 1), *feature.NewColor(1.9, 1.9, 1.9)},
		{feature.Vector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2), feature.Vector(0, 0, -1), feature.Point(0, 0, -10), feature.NewColor(1, 1, 1), *feature.NewColor(1, 1, 1)},
	}
	for _, table := range tables {
		var comp feature.Computations
		m := feature.NewMaterial()
		position := feature.Point(0, 0, 0)
		light := feature.NewLight()
		*light = light.PointLight(*table.position, *table.color)
		comp.Point = *position
		comp.Normal = *table.normal
		comp.Eye = *table.eye
		ansc := m.Lighting(*light, comp)
		if ansc != table.ansc {
			t.Errorf("You are wrong %v", table.ansc)
		}
	}
}
