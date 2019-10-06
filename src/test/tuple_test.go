package test

import (
	"feature"
	"math"
	"testing"
)

//TestTuple1 tests to see if the GetTuple function works for feature tuple
func TestTuple1(t *testing.T) {
	tables := []struct {
		v       feature.Tuple
		x, y, z float64
		point   bool
	}{
		{feature.Tuple{4.3, -4.2, 3.1, 1.0}, 4.3, -4.2, 3.1, true},
		{feature.Tuple{4.5, -42, 310, 0.0}, 4.5, -42, 310, false},
		{feature.Tuple{4.5, -42, 310, 0.0}, 4.5, -42, 310, false},
	}
	for _, table := range tables {
		x, y, z, typeOfTuple := table.v.GetTuple()
		if x != table.x || y != table.y || z != table.z || typeOfTuple != table.point {
			t.Errorf("Error Input")
		}
	}
}

//TestTuple2 tests to see if the function Point works as the way we want
func TestTuple2(t *testing.T) {
	tables := []struct {
		x, y, z float64
		ans     feature.Tuple
	}{
		{4.3, -4.2, 3.1, feature.Tuple{4.3, -4.2, 3.1, 1.0}},
	}
	for _, table := range tables {
		point := feature.Point(table.x, table.y, table.z)
		if *point != table.ans {
			t.Error("You are wrong")
		}
	}
}

//TestTuple3 tests to see if the function Vector works as the way we want
func TestTuple3(t *testing.T) {
	tables := []struct {
		x, y, z float64
		ans     feature.Tuple
	}{
		{4.3, -4.2, 3.1, feature.Tuple{4.3, -4.2, 3.1, 0}},
	}
	for _, table := range tables {
		point := feature.Vector(table.x, table.y, table.z)
		if *point != table.ans {
			t.Error("You are wrong")
		}
	}
}

//TestTuple4 tests to see if the function Add works as the way we want
func TestTuple4(t *testing.T) {
	tables := []struct {
		tone feature.Tuple
		ttwo feature.Tuple
		ans  feature.Tuple
	}{
		{feature.Tuple{4.3, -4.2, 3.1, 0}, feature.Tuple{4.3, -4.2, 3.1, 1}, feature.Tuple{8.6, -8.4, 6.2, 1}},
	}
	for _, table := range tables {
		ans := table.tone.Add(&table.ttwo)
		if ans != table.ans {
			t.Error("You are wrong")
		}
	}
}

//TestTuple5 tests to see if the function Subtract works as the way we want
//For now if we use a vector to subtract a point we will
//get a vector unless otherwise instructed
func TestTuple5(t *testing.T) {
	tables := []struct {
		tone        feature.Tuple
		ttwo        feature.Tuple
		ans         feature.Tuple
		typeOfTuple bool
	}{
		{feature.Tuple{4.3, -4.2, 3.1, 1}, feature.Tuple{4.3, -4.2, 3.1, 0}, feature.Tuple{0, 0, 0, 1}, true},
		{feature.Tuple{4.3, -4.2, 3.1, 0}, feature.Tuple{4.3, -4.2, 3.1, 0}, feature.Tuple{0, 0, 0, 0}, false},
		{feature.Tuple{4.3, -4.2, 3.1, 1}, feature.Tuple{4.3, -4.2, 3.1, 1}, feature.Tuple{0, 0, 0, 0}, false},
		{feature.Tuple{4.3, -4.2, 3.1, 0}, feature.Tuple{4.3, -4.2, 3.1, 1}, feature.Tuple{0, 0, 0, -1}, false},
		{feature.Tuple{0, 0, 0, 0}, feature.Tuple{4.3, -4.2, 3.1, 0}, feature.Tuple{-4.3, 4.2, -3.1, 0}, false},
	}
	for _, table := range tables {
		ans, typeOfTuple := table.tone.Subtract(&table.ttwo)
		if ans != table.ans || typeOfTuple != table.typeOfTuple {
			t.Error("You are wrong")
		}
	}
}

//TestTuple6 tests to see if the function Multiply works as the way we want
func TestTuple6(t *testing.T) {
	tables := []struct {
		tone feature.Tuple
		num  float64
		ans  feature.Tuple
	}{
		{feature.Tuple{4.3, -4.2, 3.1, 1}, -1, feature.Tuple{-4.3, 4.2, -3.1, -1}},
		{feature.Tuple{4.3, -4.2, 3.1, 0}, 2, feature.Tuple{8.6, -8.4, 6.2, 0}},
	}
	for _, table := range tables {
		ans := table.tone.Multiply(table.num)
		if ans != table.ans {
			t.Errorf("You are wrong, need %f,  %f, %f, %f, but get %f, %f, %f, %f", table.ans.X, table.ans.Y, table.ans.Z, table.ans.W, ans.X, ans.Y, ans.Z, ans.W)
		}
	}
}

//TestTuple7 tests to see if the function Divide works as the way we want
func TestTuple7(t *testing.T) {
	tables := []struct {
		tone feature.Tuple
		num  float64
		ans  feature.Tuple
	}{
		{feature.Tuple{4.3, -4.2, 3.1, 1}, -1, feature.Tuple{-4.3, 4.2, -3.1, -1}},
		{feature.Tuple{4.3, -4.2, 3.1, 0}, 0, feature.Tuple{0, 0, 0, 0}},
	}
	for _, table := range tables {
		ans := table.tone.Divide(table.num)
		if ans != table.ans {
			t.Errorf("You are wrong, need %f,  %f, %f, %f, but get %f, %f, %f, %f", table.ans.X, table.ans.Y, table.ans.Z, table.ans.W, ans.X, ans.Y, ans.Z, ans.W)
		}
	}
}

//TestTuple8 tests to see if the function Magnitude works as the way we want
func TestTuple8(t *testing.T) {
	tables := []struct {
		tone     feature.Tuple
		ans      float64
		vecOrNot bool
	}{
		{feature.Tuple{4.3, -4.2, 3.1, 1}, 0, false},
		{feature.Tuple{4.3, -4.2, 3.1, 0}, 6.763135367564367, true},
	}
	for _, table := range tables {
		ans, vecOrNot := table.tone.Magnitude()
		if ans != table.ans || vecOrNot != table.vecOrNot {
			t.Errorf("You are wrong")
		}
	}
}

//TestTuple9 tests to see if the function Normalize works as the way we want
func TestTuple9(t *testing.T) {
	tables := []struct {
		tone       feature.Tuple
		ans        feature.Tuple
		normalized bool
	}{
		{feature.Tuple{4.3, -4.2, 3.1, 1}, feature.Tuple{0, 0, 0, 0}, false},
		{feature.Tuple{4.3, -4.2, 3.1, 0}, feature.Tuple{0.6357997831335106, -0.6210137416652894, 0.4583672855148565, 0}, true},
		{feature.Tuple{0, 0, 0, 0}, feature.Tuple{0, 0, 0, 0}, false},
	}
	for _, table := range tables {
		ans, normalized := table.tone.Normalize()
		if ans != table.ans || normalized != table.normalized {
			t.Errorf("You are wrong")
		}
	}
}

//TestTuple10 tests to see if the function Normalize works as the way we want
//by checking whether the magnitude of the normalized  is 1
func TestTuple10(t *testing.T) {
	tables := []struct {
		tone feature.Tuple
		mag  float64
	}{
		{feature.Tuple{4.3, -4.2, 3.1, 0}, 1},
	}
	for _, table := range tables {
		ans, normalized := table.tone.Normalize()
		if normalized {
			mag, _ := ans.Magnitude()
			if !normalized || mag != 1 {
				t.Errorf("You are wrong")
			}
		}
	}
}

//TestTuple11 tests to see if the function DotProduct works as the way we want
func TestTuple11(t *testing.T) {
	tables := []struct {
		tone   feature.Tuple
		ttwo   feature.Tuple
		ans    float64
		dotted bool
	}{
		{feature.Tuple{1, 2, 3, 0}, feature.Tuple{2, 3, 4, 0}, 20, true},
		{feature.Tuple{1, 2, 3, 1}, feature.Tuple{2, 3, 4, 0}, 0, false},
		{feature.Tuple{1, 2, 3, 1}, feature.Tuple{2, 3, 4, 1}, 0, false},
	}
	for _, table := range tables {
		ans, dotted := table.tone.DotProduct(&table.ttwo)
		if ans != table.ans || dotted != table.dotted {
			t.Errorf("You are wrong, %f", ans)
		}
	}
}

//TestTuple12 tests to see if the function MagnitudeSquared works as the way we want
func TestTuple12(t *testing.T) {
	tables := []struct {
		tone     feature.Tuple
		ans      float64
		vecOrNot bool
	}{
		{feature.Tuple{4.3, -4.2, 3.1, 1}, 0, false},
		{feature.Tuple{4.3, -4.2, 3.1, 0}, 45.739999999999995, true},
	}
	for _, table := range tables {
		ans, vecOrNot := table.tone.MagnitudeSquared()
		if ans != table.ans || vecOrNot != table.vecOrNot {
			t.Errorf("You are wrong")
		}
	}
}

//TestTuple13 tests to see whether the function RotateX,RotateY,RotateZ works as we wanted
func TestTuple13(t *testing.T) {
	tables := []struct {
		point     feature.Tuple
		r         float64
		transType string
		ansO      feature.Tuple
	}{
		{*feature.Point(0, 1, 0), (math.Pi / 4), "x", *feature.Point(0, (math.Sqrt(2) / 2), (math.Sqrt(2) / 2))},
		{*feature.Point(0, 1, 0), (math.Pi / 2), "x", *feature.Point(0, 0, 1)},
		{*feature.Point(0, 0, 1), (math.Pi / 4), "y", *feature.Point((math.Sqrt(2) / 2), 0, (math.Sqrt(2) / 2))},
		{*feature.Point(0, 0, 1), (math.Pi / 2), "y", *feature.Point(1, 0, 0)},
		{*feature.Point(0, 1, 0), (math.Pi / 4), "z", *feature.Point(-(math.Sqrt(2) / 2), (math.Sqrt(2) / 2), 0)},
		{*feature.Point(0, 1, 0), (math.Pi / 2), "z", *feature.Point(-1, 0, 0)},
	}
	for _, table := range tables {
		matrix := feature.NewMatrix(4, 4)
		if table.transType == "x" {
			matrix = feature.RotationX(table.r)
		} else if table.transType == "y" {
			matrix = feature.RotationY(table.r)
		} else if table.transType == "z" {
			matrix = feature.RotationZ(table.r)
		} else {
			t.Errorf("Wrong transform type")
		}
		errorAllowance := 0.0000001
		ans, _ := matrix.MultiplyTuple(&table.point)
		if math.Abs(ans.X-table.ansO.X) > errorAllowance || math.Abs(ans.Y-table.ansO.Y) > errorAllowance || math.Abs(ans.Z-table.ansO.Z) > errorAllowance {
			t.Errorf("Error Input %v, %v, %v, %v", table.r, table.transType, ans, table.ansO)
		}
	}
}

//TestTuple14 tests to see whether the function Shearing works as we wanted
func TestTuple14(t *testing.T) {
	tables := []struct {
		point                  feature.Tuple
		xy, xz, yx, yz, zx, zy float64
		ansO                   feature.Tuple
	}{
		{*feature.Point(2, 3, 4), 1, 0, 0, 0, 0, 0, *feature.Point(5, 3, 4)},
		{*feature.Point(2, 3, 4), 0, 1, 0, 0, 0, 0, *feature.Point(6, 3, 4)},
		{*feature.Point(2, 3, 4), 0, 0, 1, 0, 0, 0, *feature.Point(2, 5, 4)},
		{*feature.Point(2, 3, 4), 0, 0, 0, 1, 0, 0, *feature.Point(2, 7, 4)},
		{*feature.Point(2, 3, 4), 0, 0, 0, 0, 1, 0, *feature.Point(2, 3, 6)},
		{*feature.Point(2, 3, 4), 0, 0, 0, 0, 0, 1, *feature.Point(2, 3, 7)},
	}
	for _, table := range tables {
		matrix := feature.Shearing(table.xy, table.xz, table.yx, table.yz, table.zx, table.zy)
		errorAllowance := 0.0000001
		ans, _ := matrix.MultiplyTuple(&table.point)
		if math.Abs(ans.X-table.ansO.X) > errorAllowance || math.Abs(ans.Y-table.ansO.Y) > errorAllowance || math.Abs(ans.Z-table.ansO.Z) > errorAllowance {
			t.Errorf("Error Input %v, %v", ans, table.ansO)
		}
	}
}

//TestTuple15 tests to see whether the function Reflect works as we wanted
func TestTuple15(t *testing.T) {
	tables := []struct {
		incidentRay feature.Tuple
		normal      feature.Tuple
		ans         feature.Tuple
	}{
		{*feature.Vector(0, -1, 0), *feature.Vector(math.Sqrt(2)/2, math.Sqrt(2)/2, 0), *feature.Vector(1, 0, 0)},
		{*feature.Vector(1, -1, 0), *feature.Vector(0, 1, 0), *feature.Vector(1, 1, 0)},
	}
	for _, table := range tables {
		errorAllowance := 0.000000001
		ans, _ := table.incidentRay.Reflect(&table.normal)
		if math.Abs(ans.X-table.ans.X) > errorAllowance || math.Abs(ans.Y-table.ans.Y) > errorAllowance || math.Abs(ans.Z-table.ans.Z) > errorAllowance {
			t.Errorf("Error Input %v, %v", ans, table.ans)
		}
	}
}
