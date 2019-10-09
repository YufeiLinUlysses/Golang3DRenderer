package test

import (
	"feature"
	"math"
	"testing"
)

//TestObject1 tests to see if the ViewTransformation  function works for feature Object
func TestObject1(t *testing.T) {
	tables := []struct {
		from, to, up *feature.Tuple
		ans          []float64
	}{
		{feature.Point(0, 0, 0), feature.Point(0, 0, -1), feature.Vector(0, 1, 0), []float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}},
		{feature.Point(0, 0, 0), feature.Point(0, 0, 1), feature.Vector(0, 1, 0), []float64{-1, 0, 0, 0, 0, 1, 0, 0, 0, 0, -1, 0, 0, 0, 0, 1}},
		{feature.Point(0, 0, 8), feature.Point(0, 0, 0), feature.Vector(0, 1, 0), []float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, -8, 0, 0, 0, 1}},
		{feature.Point(1, 3, 2), feature.Point(4, -2, 8), feature.Vector(1, 1, 0), []float64{-0.50709, 0.50709, 0.67612, -2.36643, 0.76772, 0.60609, 0.12122, -2.82843, -0.35857, 0.59761, -0.71714, 0, 0, 0, 0, 1}},
	}
	for _, table := range tables {
		m := feature.ViewTransformation(*table.from, *table.to, *table.up)
		count := 0
		errorAllowance := 0.0001
		for i, row := range m.Matrix {
			for j := range row {
				tempAns := m.GetValueAt(j, i)
				if math.Abs(tempAns-table.ans[count]) > errorAllowance {
					t.Errorf("You are wrong %v, %v,%v", table.from, tempAns, table.ans[count])
					break
				}
				count++
			}
		}
	}
}
