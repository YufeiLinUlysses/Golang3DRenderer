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
