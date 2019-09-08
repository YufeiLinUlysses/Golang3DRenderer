package test

import (
	"class"
	"testing"
)

//TestRay1 tests to see if the GetRay function works for class tuple
func TestRay1(t *testing.T) {
	tables := []struct {
		oriX, oriY, oriZ, dirX, dirY, dirZ float64
		ansOri                             class.Tuple
		ansDir                             class.Tuple
	}{
		{1, 2, 3, 1, 2, 3, class.Tuple{1, 2, 3, 1}, class.Tuple{1, 2, 3, 0}},
	}
	for _, table := range tables {
		r := class.NewRay(table.oriX, table.oriY, table.oriZ, table.dirX, table.dirY, table.dirZ)
		ori, dir := r.GetRay()
		if ori != table.ansOri || dir != table.ansDir {
			t.Errorf("Error Input %v,%v", ori, dir)
		}
	}
}

func TestRay2(t *testing.T) {
	tables := []struct {
		oriX, oriY, oriZ, dirX, dirY, dirZ float64
		dist                               float64
		ans                                class.Tuple
	}{
		{1, 2, 3, 1, 2, 3, 3, class.Tuple{4, 8, 12, 1}},
	}
	for _, table := range tables {
		r := class.NewRay(table.oriX, table.oriY, table.oriZ, table.dirX, table.dirY, table.dirZ)
		ans := r.Position(table.dist)
		if ans != table.ans {
			t.Errorf("Error Input %v", ans)
		}
	}
}
