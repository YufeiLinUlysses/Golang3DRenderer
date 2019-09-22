package test

import (
	"class"
	"math"
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
		{4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 0, 0, 0, 1}, 4, 4, []float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 0, 0, 0, 1}},
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

//TestMatrix4 tests to see if the MultiplyTuple function works for class Matrix
func TestMatrix4(t *testing.T) {
	tables := []struct {
		w, h     int
		value    []float64
		tuple    class.Tuple
		ansTuple class.Tuple
	}{
		{4, 4, []float64{1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1}, *class.NewTuple(1, 2, 3, 1), *class.NewTuple(18, 24, 33, 1)},
		{4, 4, []float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}, *class.NewTuple(1, 2, 3, 4), *class.NewTuple(1, 2, 3, 4)},
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
		ansM, _ := m.MultiplyTuple(&table.tuple)
		if *ansM != table.ansTuple {
			t.Errorf("You are wrong")
		}
	}
}

//TestMatrix5 tests to see if the Determinant function works for class Matrix
func TestMatrix5(t *testing.T) {
	tables := []struct {
		w, h       int
		value      []float64
		ans        float64
		invertible bool
	}{
		{2, 2, []float64{1, 5, -3, 2}, 17, true},
		{3, 3, []float64{1, 2, 6, -5, 8, -4, 2, 6, 4}, -196, true},
		{4, 4, []float64{-2, -8, 3, 5, -3, 1, 7, 3, 1, 2, -9, 6, 0, 0, 0, 1}, 185, true},
		{4, 4, []float64{6, 4, 4, 4, 5, 5, 7, 6, 4, -9, 3, -7, 0, 0, 0, 1}, 260, true},
		{4, 4, []float64{-4, 2, 0, -3, 9, 6, 0, 6, 0, -5, 0, -5, 0, 0, 0, 1}, 0, false},
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
		ansM, invertible := m.Determinant()
		if ansM != table.ans && invertible == table.invertible {
			t.Errorf("You are wrong %v", ansM)
		}
	}
}

//TestMatrix6 tests to see if the SubMatrix function works for class Matrix
func TestMatrix6(t *testing.T) {
	tables := []struct {
		w, h  int
		value []float64
		ans   []float64
	}{
		{3, 3, []float64{1, 5, 0, -3, 2, 7, 0, 6, -3}, []float64{-3, 2, 0, 6}},
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
		ansM := m.SubMatrix(0, 2)
		count = 0
		for i, row := range ansM.Matrix {
			for j := range row {
				tempAns := ansM.GetValueAt(j, i)
				if tempAns != table.ans[count] {
					t.Errorf("You are wrong: %f should be: %f at place %d,%d", tempAns, table.ans[count], i, j)
				}
				count++
			}
		}
	}
}

//TestMatrix7 tests to see if the GetInverse function works for class Matrix
func TestMatrix7(t *testing.T) {
	tables := []struct {
		w, h  int
		value []float64
		ans   []float64
	}{
		{4, 4, []float64{-5, 2, 6, -8, 1, -5, 1, 8, 7, 7, -6, -7, 0, 0, 0, 1}, []float64{0.141104, 0.331288, 0.196319, -0.147239, 0.079755, -0.073620, 0.067485, 1.699387, 0.257669, 0.300613, 0.141104, 0.644172, 0.0, 0.0, 0.0, 1.0}},
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
		deter, _ := m.Determinant()
		inv := m.GetInverse(deter)
		count = 0
		for i, row := range inv.Matrix {
			for j := range row {
				tempAns := inv.GetValueAt(j, i)
				if math.Round(tempAns*1000) != math.Round(table.ans[count]*1000) {
					t.Errorf("You are wrong: %f should be: %f at place %d,%d", tempAns, table.ans[count], i, j)
				}
				count++
			}
		}
	}
}

//TestMatrix8 tests to see if a process works for class Matrix
func TestMatrix8(t *testing.T) {
	tables := []struct {
		w1, h1   int
		value1   []float64
		w2, h2   int
		value2   []float64
		ansValue []float64
	}{
		{4, 4, []float64{3,-9,7,3,3,-8,2,-9,-4,4,4,1,0,0,0,1}, 
		 4, 4, []float64{8,2,2,2,3,-1,7,0,7,0,5,4,0,0,0,1},
		 []float64{3,-9,7,3,3,-8,2,-9,-4,4,4,1,0,0,0,1}},
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
		mult,_ := m.Multiply(m2)
		deter,_ := m2.Determinant()
		rev,_:=mult.Multiply(m2.GetInverse(deter))
		for i, row := range rev.Matrix {
			for j := range row {
				tempAns := rev.GetValueAt(j, i)
				if math.Round(tempAns*1000) != math.Round(table.ansValue[count]*1000) {
					t.Errorf("You are wrong: %f should be: %f at place %d,%d", tempAns, table.ansValue[count], i, j)
				}
				count++
			}
		}
	}
}
