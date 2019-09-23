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
		{1, 2, 3, 4, 5, 6, class.Tuple{1, 2, 3, 1}, class.Tuple{0.4558423058385518, 0.5698028822981898, 0.6837634587578276, 0}},
	}
	for _, table := range tables {
		r := class.NewRay(table.oriX, table.oriY, table.oriZ, table.dirX, table.dirY, table.dirZ)
		ori, dir := r.GetRay()
		if ori != table.ansOri || dir != table.ansDir {
			t.Errorf("Error Input %v,%v", ori, dir)
		}
	}
}

//TestRay2 tests to see whether the function Position works as we wanted
func TestRay2(t *testing.T) {
	tables := []struct {
		oriX, oriY, oriZ, dirX, dirY, dirZ float64
		dist                               float64
		ans                                class.Tuple
	}{
		{2, 3, 4, 1, 0, 0, 0, *class.Point(2, 3, 4)},
		{2, 3, 4, 1, 0, 0, 1, *class.Point(3, 3, 4)},
		{2, 3, 4, 1, 0, 0, -1, *class.Point(1, 3, 4)},
		{2, 3, 4, 1, 0, 0, 2.5, *class.Point(4.5, 3, 4)},
	}
	for _, table := range tables {
		r := class.NewRay(table.oriX, table.oriY, table.oriZ, table.dirX, table.dirY, table.dirZ)
		ans := r.Position(table.dist)
		if ans != table.ans {
			t.Errorf("Error Input %v", ans)
		}
	}
}

//TestRay3 tests to see whether the function Translate works as we wanted
func TestRay3(t *testing.T) {
	tables := []struct {
		oriX, oriY, oriZ, dirX, dirY, dirZ float64
		xInc, yInc, zInc                   float64
		ansO                               class.Tuple
		ansD                               class.Tuple
	}{
		{1, 2, 3, 0, 1, 0, 3, 4, 5, *class.Point(4, 6, 8), *class.Vector(0, 1, 0)},
	}
	for _, table := range tables {
		r := class.NewRay(table.oriX, table.oriY, table.oriZ, table.dirX, table.dirY, table.dirZ)
		ans := r.Translate(table.xInc, table.yInc, table.zInc)
		if ans.Origin != table.ansO || ans.Direction != table.ansD {
			t.Errorf("Error Input %v", ans)
		}
	}
}

//TestRay4 tests to see whether the function Scale works as we wanted
func TestRay4(t *testing.T) {
	tables := []struct {
		oriX, oriY, oriZ, dirX, dirY, dirZ float64
		xInc, yInc, zInc                   float64
		ansO                               class.Tuple
		ansD                               class.Tuple
	}{
		{1, 2, 3, 0, 1, 0, 2, 3, 4, *class.Point(2, 6, 12), *class.Vector(0, 3, 0)},
	}
	for _, table := range tables {
		r := class.NewRay(table.oriX, table.oriY, table.oriZ, table.dirX, table.dirY, table.dirZ)
		ans := r.Scale(table.xInc, table.yInc, table.zInc)
		if ans.Origin != table.ansO || ans.Direction != table.ansD {
			t.Errorf("Error Input %v", ans)
		}
	}
}
