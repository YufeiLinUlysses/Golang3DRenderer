package test

import (
	"feature"
	"math"
	"testing"
)

//TestComputation1 tests to see if the Schlick function works for feature Computations
func TestComputation1(t *testing.T) {
	shape := feature.NewSphere()
	shape.Mat.Transparency = 1
	shape.Mat.Refractivity = 1.5
	tables := []struct {
		r      feature.Ray
		i1, i2 float64
		index  int
		ans    float64
	}{
		{*feature.NewRay(*feature.Point(0, 0, math.Sqrt(2)/2), *feature.Vector(0, 1, 0)), -math.Sqrt(2) / 2, math.Sqrt(2) / 2, 1, 1},
		{*feature.NewRay(*feature.Point(0, 0, 0), *feature.Vector(0, 1, 0)), -1, 1, 1, 0.04000000000000001},
		{*feature.NewRay(*feature.Point(0, 0.99, -2), *feature.Vector(0, 0, 1)), 1.8589, 1.8589, 0, 0.4887308101221217},
	}

	for _, table := range tables {
		r := table.r
		i1 := feature.NewIntersection(table.i1, r, shape)
		i2 := feature.NewIntersection(table.i2, r, shape)
		var xs []feature.Intersection
		if table.i1 != table.i2 {
			xs = append(xs, *i1, *i2)
		} else {
			xs = append(xs, *i1)
		}
		comps := xs[table.index].PrepareComputation(&r, xs)
		if comps.Schlick() != table.ans {
			t.Errorf("Error input %v", comps.Schlick())
		}
	}
}
