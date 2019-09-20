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
		{2, 2, []float64{-3, 5, 1, -2}, []float64{-3, 5, 1, -2}},
		{3, 3, []float64{-3, 5, 0, 1, -2, -7, 0, 1, 1}, []float64{-3, 5, 0, 1, -2, -7, 0, 1, 1}},
	}
	for _, table := range tables {
		count := 0
		m := class.NewMatrix(table.w, table.h)
		for i, row := range m.Matrix {
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

//TestMatrix2 tests to see if the EqualTo function works for class Matrix
func TestMatrix2(t *testing.T) {
	tables := []struct {
		w1, h1 int
		value1 []float64
		w2, h2 int
		value2 []float64
		ans    bool
	}{
		{4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 0, 0, 0, 1}, 4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 0, 0, 0, 1}, true},
		{4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 0, 0, 0, 1}, 4, 4, []float64{2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 0, 0, 0, 1}, false},
	}
	for _, table := range tables {
		count := 0
		m := class.NewMatrix(table.w1, table.h1)
		for i, row := range m.Matrix {
			for j := range row {
				m = m.Assign(j, i, table.value1[count])
				count++
			}
		}
		count = 0
		m2 := class.NewMatrix(table.w2, table.h2)
		for i, row := range m.Matrix {
			for j := range row {
				m2 = m2.Assign(j, i, table.value2[count])
				count++
			}
		}
		if m.EqualTo(*m2) != table.ans {
			t.Errorf("There is a problem %v, %v", m.EqualTo(*m2), table.ans)
		}
	}
}

//TestMatrix3 tests to see if the Multiply function works for class Matrix
func TestMatrix3(t *testing.T) {
	tables := []struct {
		w1, h1   int
		value1   []float64
		w2, h2   int
		value2   []float64
		ansValue []float64
	}{
		{4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 0, 0, 0, 1}, 4, 4, []float64{-2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 0, 0, 0, 1}, []float64{16, 14, 22, 20, 36, 38, 58, 52, 34, 46, 68, 60, 0, 0, 0, 1}},
	}
	for _, table := range tables {
		count := 0
		m := class.NewMatrix(table.w1, table.h1)
		for i, row := range m.Matrix {
			for j := range row {
				m = m.Assign(j, i, table.value1[count])
				count++
			}
		}
		count = 0
		m2 := class.NewMatrix(table.w2, table.h2)
		for i, row := range m2.Matrix {
			for j := range row {
				m2 = m2.Assign(j, i, table.value2[count])
				count++
			}
		}
		count = 0
		ansM, _ := m.Multiply(m2)
		for i, row := range ansM.Matrix {
			for j := range row {
				tempAns := ansM.GetValueAt(j, i)
				if tempAns != table.ansValue[count] {
					t.Errorf("You are wrong: %f should be: %f at place %d,%d", tempAns, table.ansValue[count], i, j)
				}
				count++
			}
		}
	}
}
