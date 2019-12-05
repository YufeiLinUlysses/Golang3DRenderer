package test

/*All equations are in the form  coef[0] + coef[1]*x + coef[2]*x^2 + coef[3]*x^3 + coef[4]*x^4 +...= 0*/
import (
	"feature"
	"testing"
)

//TestFindRoots1 tests to see if the SolveQuadratic  function works for feature findroots
func TestFindRoots1(t *testing.T) {
	tables := []struct {
		coef    []float64
		hasRoot bool
		ans     []float64
	}{
		{[]float64{1, 2, 3}, false, []float64{}},
		{[]float64{1, 2, 1}, true, []float64{-1}},
		{[]float64{1, 5, 1}, true, []float64{-0.20871215252208009, -4.7912878474779195}},
	}
	for _, table := range tables {
		root := feature.SolveQuadratic(table.coef)
		if root.Count == len(table.ans) {
			for i := 0; i < root.Count; i++ {
				if root.Ans[i] != table.ans[i] {
					t.Errorf("Error Input")
				}
			}
		}
	}
}

//TestFindRoots2 tests to see if the SolveCubic  function works for feature findroots
func TestFindRoots2(t *testing.T) {
	tables := []struct {
		coef    []float64
		hasRoot bool
		ans     []float64
	}{
		{[]float64{1, 3, 3, 1}, true, []float64{-1}},
		{[]float64{1, 3, 4, 1}, true, []float64{-3.1478990357047874}},
		{[]float64{-6, 11, -6, 1}, true, []float64{3, 2, 1}},
	}
	for _, table := range tables {
		root := feature.SolveCubic(table.coef)
		if root.Count == len(table.ans) {
			for i := 0; i < root.Count; i++ {
				if root.Ans[i] != table.ans[i] {
					t.Errorf("Error Input")
				}
			}
		}
	}
}

//TestFindRoots3 tests to see if the SolveQuartic  function works for feature findroots
func TestFindRoots3(t *testing.T) {
	tables := []struct {
		coef    []float64
		hasRoot bool
		ans     []float64
	}{
		{[]float64{5, 10, 10, 5, 1}, false, []float64{}},
		{[]float64{1, 4, 6, 4, 1}, true, []float64{-1}},
		{[]float64{1, 4, 6, 5, 1}, true, []float64{-0.4502995220980296, -3.6296581267545345}},
	}
	for _, table := range tables {
		root := feature.SolveQuartic(table.coef)
		if root.Count == len(table.ans) {
			for i := 0; i < root.Count; i++ {
				if root.Ans[i] != table.ans[i] {
					t.Errorf("Error Input")
				}
			}
		}
	}
}
