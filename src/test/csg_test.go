package test

import (
	"feature"
	"testing"
)

//TestColor1 tests to see if the IntersectionAllowed  function works for feature CSG
func TestCSG1(t *testing.T) {
	tables := []struct {
		op             string
		lhit, inl, inr bool
		ans            bool
	}{
		{"union", true, true, true, false},
		{"union", true, true, false, true},
		{"union", true, false, true, false},
		{"union", true, false, false, true},
		{"union", false, true, true, false},
		{"union", false, true, false, false},
		{"union", false, false, true, true},
		{"union", false, false, false, true},

		{"intersection", true, true, true, true},
		{"intersection", true, true, false, false},
		{"intersection", true, false, true, true},
		{"intersection", true, false, false, false},
		{"intersection", false, true, true, true},
		{"intersection", false, true, false, true},
		{"intersection", false, false, true, false},
		{"intersection", false, false, false, false},

		{"difference", true, true, true, false},
		{"difference", true, true, false, true},
		{"difference", true, false, true, false},
		{"difference", true, false, false, true},
		{"difference", false, true, true, true},
		{"difference", false, true, false, true},
		{"difference", false, false, true, false},
		{"difference", false, false, false, false},
	}
	for _, table := range tables {
		s := feature.NewSphere()
		c := feature.NewCube()
		csg := feature.NewCSG("", s, c)
		ans := csg.IntersectionAllowed(table.op,table.lhit,table.inl,table.inr)
		if ans != table.ans{
			t.Errorf("Error Input")
		}
	}
}
