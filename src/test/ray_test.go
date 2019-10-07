package test

import (
	"feature"
	"testing"
)

//TestRay1 tests to see if the GetRay function works for feature tuple
func TestRay1(t *testing.T) {
	tables := []struct {
		origin, direction feature.Tuple
		ansOri            feature.Tuple
		ansDir            feature.Tuple
	}{
		{*feature.Point(1, 2, 3), *feature.Vector(4, 5, 6), feature.Tuple{1, 2, 3, 1}, feature.Tuple{0.4558423058385518, 0.5698028822981898, 0.6837634587578276, 0}},
	}
	for _, table := range tables {
		r := feature.NewRay(table.origin, table.direction)
		ori, dir := r.GetRay()
		if ori != table.ansOri || dir != table.ansDir {
			t.Errorf("Error Input %v,%v", ori, dir)
		}
	}
}

//TestRay2 tests to see whether the function Position works as we wanted
func TestRay2(t *testing.T) {
	tables := []struct {
		origin, direction feature.Tuple
		dist              float64
		ans               feature.Tuple
	}{
		{*feature.Point(2, 3, 4), *feature.Vector(1, 0, 0), 0, *feature.Point(2, 3, 4)},
		{*feature.Point(2, 3, 4), *feature.Vector(1, 0, 0), 1, *feature.Point(3, 3, 4)},
		{*feature.Point(2, 3, 4), *feature.Vector(1, 0, 0), -1, *feature.Point(1, 3, 4)},
		{*feature.Point(2, 3, 4), *feature.Vector(1, 0, 0), 2.5, *feature.Point(4.5, 3, 4)},
	}
	for _, table := range tables {
		r := feature.NewRay(table.origin, table.direction)
		ans := r.Position(table.dist)
		if ans != table.ans {
			t.Errorf("Error Input %v", ans)
		}
	}
}

//TestRay3 tests to see whether the function Transform works as we wanted
func TestRay3(t *testing.T) {
	tables := []struct {
		origin, direction feature.Tuple
		xInc, yInc, zInc  float64
		transType         string
		ansO              feature.Tuple
		ansD              feature.Tuple
	}{
		{*feature.Point(1, 2, 3), *feature.Vector(0, 1, 0), 3, 4, 5, "translate", *feature.Point(4, 6, 8), *feature.Vector(0, 1, 0)},
		{*feature.Point(1, 2, 3), *feature.Vector(0, 1, 0), 2, 3, 4, "scale", *feature.Point(2, 6, 12), *feature.Vector(0, 3, 0)},
	}
	for _, table := range tables {
		r := feature.NewRay(table.origin, table.direction)
		matrix := feature.NewMatrix(4, 4)
		if table.transType == "translate" {
			matrix = feature.Translate(table.xInc, table.yInc, table.zInc)
		} else if table.transType == "scale" {
			matrix = feature.Scale(table.xInc, table.yInc, table.zInc)
		} else {
			t.Errorf("Wrong transform type")
		}
		ans := r.Transform(matrix)
		if ans.Origin != table.ansO || ans.Direction != table.ansD {
			t.Errorf("Error Input %v, %v", ans, table.ansO)
		}
	}
}
