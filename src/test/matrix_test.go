package test

import (
	"class"
	"testing"
)

//TestMatrix1 tests to see if the Assign and GetValueAt functions work for class Matrix
func TestMatrix1(t *testing.T) {
	tables := []struct {
		w, h     int
		value    []float64
		ansValue []float64
	}{
		{4, 4, []float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 0, 0, 0, 1}, []float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 0, 0, 0, 1}},
	}
	for _, table := range tables {
		count := 0
		m := class.NewMatrix(table.w, table.h)
		for i, row := range m .Matrix{
			for j := range row {
				m = m.Assign(j, i, table.value[count])
				count++
			}
		}
		count = 0
		for i, row := range m.Matrix {
			for j := range row {
				tempAns := m.GetValueAt(j, i)
				if tempAns != table.ansValue[count] {
					t.Errorf("You are wrong: %f should be: %f at place %d,%d", tempAns, table.ansValue[count], i, j)
				}
				count++
			}
		}
	}
}
