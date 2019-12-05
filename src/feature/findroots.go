package feature

/*This is a translation  from c to golang of
 *https://github.com/erich666/GraphicsGems/blob/master/gems/Roots3And4.c
 *to solve quartic equation*/

import (
	"math"
)

/*General quadratic equation looks like  coef[0] + coef[1]*x + coef[2]*x^2 = 0
 *General quadratic equation looks like  coef[0] + coef[1]*x + coef[2]*x^2 + coef[3]*x^3 = 0
 *General quartic equation looks like  coef[0] + coef[1]*x + coef[2]*x^2 + coef[3]*x^3 + coef[4]*x^4 = 0*/

/*Roots contains all necessary roots*/
type Roots struct {
	Count   int
	Ans     []float64
	HasRoot bool
}

/*unique is a helper function to get rid of duplicated elements in a slice*/
func unique(intSlice []float64) []float64 {
	keys := make(map[float64]bool)
	list := []float64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

/*IsZero checks whether a number is zero*/
func IsZero(num float64) bool {
	if math.Pow10(-9) > num && -math.Pow10(-9) < num {
		return true
	}
	return false
}

/*NewRoot establishes a new root*/
func NewRoot() *Roots {
	r := &Roots{
		Count:   0,
		Ans:     make([]float64, 0),
		HasRoot: false,
	}
	return r
}

/*SolveQuadratic solves quadratic formula*/
func SolveQuadratic(coefs []float64) *Roots {
	root := NewRoot()
	/*coef[0] + coef[1]*x + coef[2]*x^2 = 0
	 *reduces to normal form
	 *x^2+px+q = 0*/
	p := coefs[1] / (2 * coefs[2])
	q := coefs[0] / coefs[2]
	delta := p*p - q
	if IsZero(delta) {
		root.Count = 1
		root.Ans = append(root.Ans, -p)
		root.HasRoot = true
	} else if delta > 0 {
		sqrtdelta := math.Sqrt(delta)
		ans1 := sqrtdelta - p
		ans2 := -sqrtdelta - p
		root.Count = 2
		root.Ans = append(root.Ans, ans1, ans2)
	}
	return root
}

/*SolveCubic solves cubic formula*/
func SolveCubic(coefs []float64) *Roots {
	root := NewRoot()
	/*coef[0] + coef[1]*x + coef[2]*x^2 + coef[3]*x^3 = 0
	 *reduces to normal form
	 *x^3 + ax^2 + bx + c = 0*/
	a := coefs[2] / coefs[3]
	b := coefs[1] / coefs[3]
	c := coefs[0] / coefs[3]

	/* substitute x = y - a/3 to eliminate quadric term:
	 *x^3 +px + q = 0 */
	aSquare := a * a
	p := 1.0 / 3 * (-1.0/3*aSquare + b)
	q := 1.0 / 2 * (2.0/27*a*aSquare - 1.0/3*a*b + c)

	/* use Cardano's formula */
	pCube := math.Pow(p, 3)
	delta := q*q + pCube
	if IsZero(delta) {
		if IsZero(q) {
			root.Count = 1
			root.Ans = append(root.Ans, 0)
			root.HasRoot = true
		} else {
			u := math.Pow(-q, 1.0/3)
			ans1 := 2 * u
			ans2 := -u
			root.Count = 2
			root.Ans = append(root.Ans, ans1, ans2)
			root.HasRoot = true
		}
	} else if delta < 0 {
		phi := 1.0 / 3 * math.Acos(-q/math.Sqrt(-pCube))
		t := 2 * math.Sqrt(-p)
		ans1 := t * math.Cos(phi)
		ans2 := -t * math.Cos(phi+math.Pi/3)
		ans3 := -t * math.Cos(phi-math.Pi/3)
		root.Count = 3
		root.Ans = append(root.Ans, ans1, ans2, ans3)
		root.HasRoot = true
	} else {
		deltaSqrt := math.Sqrt(delta)
		u := math.Cbrt(deltaSqrt - q)
		v := -math.Cbrt(deltaSqrt + q)
		ans1 := u + v
		root.Count = 1
		root.Ans = append(root.Ans, ans1)
		root.HasRoot = true
	}
	for i := 0; i < root.Count; i++ {
		root.Ans[i] += -a / 3
	}
	return root
}

/*SolveQuartic solves quartic formula*/
func SolveQuartic(coefs []float64) *Roots {
	root := NewRoot()
	co := make([]float64, 4)
	/*coef[0] + coef[1]*x + coef[2]*x^2 + coef[3]*x^3+coef[4] = 0
	 *reduces to normal form
	 *x^4 + ax^3 + bx^2 + cx + d = 0*/
	a := coefs[3] / coefs[4]
	b := coefs[2] / coefs[4]
	c := coefs[1] / coefs[4]
	d := coefs[0] / coefs[4]

	/*  substitute x = y - A/4 to eliminate cubic term:
	x^4 + px^2 + qx + r = 0 */
	aSquare := a * a
	p := -3.0/8*aSquare + b
	q := 1.0/8*aSquare*a - 1.0/2*a*b + c
	r := -3.0/256*aSquare*aSquare + 1.0/16*aSquare*b - 1.0/4*a*c + d
	if IsZero(r) {
		co[0] = q
		co[1] = p
		co[2] = 0
		co[3] = 1
		root = SolveCubic(co)
		root.Count = root.Count + 1
		root.Ans = append(root.Ans, 0)
		root.HasRoot = true
	} else {
		co[0] = 1.0/2*r*p - 1.0/8*q*q
		co[1] = -r
		co[2] = -1.0 / 2 * p
		co[3] = 1
		tempRoot := SolveCubic(co)
		z := tempRoot.Ans[0]
		u := z*z - r
		v := 2*z - p
		if IsZero(u) {
			u = 0
		} else if u > 0 {
			u = math.Sqrt(u)
		} else {
			return root
		}
		if IsZero(v) {
			v = 0
		} else if v > 0 {
			v = math.Sqrt(v)
		} else {
			return root
		}
		co[0] = z - u
		if q < 0 {
			co[1] = -v
		} else {
			co[1] = v
		}
		co[2] = 1
		tempRoot1 := SolveQuadratic(co)
		co[0] = z + u
		if q < 0 {
			co[1] = v
		} else {
			co[1] = -v
		}
		co[2] = 1
		tempRoot2 := SolveQuadratic(co)
		root.Ans = append(root.Ans, tempRoot1.Ans...)
		root.Ans = append(root.Ans, tempRoot2.Ans...)
		root.Count = tempRoot1.Count + tempRoot2.Count
		if root.Count > 0 {
			root.HasRoot = true
		}
	}
	for i := 0; i < root.Count; i++ {
		root.Ans[i] += -a / 4
		if IsZero(root.Ans[i]) {
			root.Ans[i] = 0
		}
	}
	root.Ans = unique(root.Ans)
	root.Count = len(root.Ans)
	return root
}
