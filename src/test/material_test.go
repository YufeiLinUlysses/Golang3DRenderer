package test

import (
	"class"
	"testing"
)

//TestLight1 tests to see if the PointLight function works for class Light
func TestLight1(t *testing.T) {
	tables := []struct {
		m    class.Material	
		ansc class.Color
		ansd float64
	}{
		{*class.NewMaterial(), *class.NewColor(1,1,1),1},
	}
	for _, table := range tables {
		ansc, ansd := table.m.GetMaterial()
		if ansc != table.ansc || ansd != table.ansd{
			t.Errorf("You are wrong")
		}
	}
}
