package feature

import (
	"math"
)

//Matrix type
type Matrix struct {
	Width, Height int
	Matrix        [][]float64
}

//NewMatrix sets up a new matrix instance
func NewMatrix(w, h int) *Matrix {
	m := make([][]float64, h)
	for i := 0; i < h; i++ {
		m[i] = make([]float64, w)
	}

	matrix := &Matrix{
		Width:  w,
		Height: h,
		Matrix: m,
	}
	return matrix
}

//Assign assigns value for a specific value on an element in a matrix
func (m *Matrix) Assign(cl, rw int, value float64) *Matrix {
	m.Matrix[rw][cl] = value
	return m
}

//GetValueAt returns the value of a specific place
func (m *Matrix) GetValueAt(cl, rw int) float64 {
	return m.Matrix[rw][cl]
}

//EqualTo determines whether two matrix are the same
func (m *Matrix) EqualTo(m2 Matrix) bool {
	if m.Width == m2.Width && m.Height == m2.Height {
		for i, row := range m.Matrix {
			for j := range row {
				mVal := m.GetValueAt(j, i)
				m2Val := m2.GetValueAt(j, i)
				if m2Val != mVal {
					return false
				}
			}
		}
		return true
	}
	return false
}

//GetIdentity returns the identity matrix of the size of the matrix
func (m *Matrix) GetIdentity() (identity *Matrix, square bool) {
	if m.Width != m.Height {
		return identity, false
	}
	identity = NewMatrix(m.Width, m.Height)
	for i := 0; i < m.Height; i++ {
		identity.Assign(i, i, 1)
	}
	return identity, true
}

//Multiply multiplies two matrices
func (m *Matrix) Multiply(m2 *Matrix) (ansM *Matrix, multiplied bool) {
	ansM = NewMatrix(m.Height, m2.Height)
	if m.Width != m2.Height {
		return ansM, false
	}
	for i := range m.Matrix {
		for j := range m2.Matrix {
			Value := float64(0)
			for k := range ansM.Matrix {
				Value += m.GetValueAt(k, i) * m2.GetValueAt(j, k)
			}
			ansM.Assign(j, i, Value)
		}
	}
	return ansM, true
}

//MultiplyTuple multiplies tuple
func (m *Matrix) MultiplyTuple(t *Tuple) (ansT *Tuple, tuple bool) {
	tup := []float64{t.X, t.Y, t.Z, t.W}
	var ans []float64
	if m.Height != 4 {
		return ansT, false
	}
	for i := 0; i < 4; i++ {
		temp := float64(0)
		for j := 0; j < 4; j++ {
			temp += m.GetValueAt(j, i) * tup[j]
		}
		ans = append(ans, temp)
	}
	ansT = NewTuple(ans[0], ans[1], ans[2], ans[3])
	return ansT, true
}

//MultiplyScalar mutlipliese the matrix to a scalar
func (m *Matrix) MultiplyScalar(scalar float64) *Matrix {
	for i, row := range m.Matrix {
		for j := range row {
			currValue := m.GetValueAt(j, i)
			m.Assign(j, i, scalar*currValue)
		}
	}
	return m
}

//SubMatrix gets the cofactor of a matrix at a certain location
func (m *Matrix) SubMatrix(cl, rw int) *Matrix {
	var value []float64
	ans := NewMatrix(m.Width-1, m.Height-1)
	for i := range m.Matrix {
		if i != cl {
			for j := range m.Matrix {
				if j != rw {
					value = append(value, m.Matrix[i][j])
				}
			}
		}
	}
	count := 0
	for i := 0; i < m.Width-1; i++ {
		for j := 0; j < m.Width-1; j++ {
			ans.Assign(j, i, value[count])
			count++
		}
	}
	return ans
}

//Determinant gets the determinant of a 2x2 square matrix
func (m *Matrix) Determinant() (ans float64, invertible bool) {
	var cl int
	if m.Width == 2 {
		ans = m.Matrix[0][0]*m.Matrix[1][1] - m.Matrix[0][1]*m.Matrix[1][0]
		return ans, true
	}
	if m.Width == 4 {
		cl = 3
	}
	for i := 0; i < len(m.Matrix); i++ {
		if m.GetValueAt(i, cl) == 0 {
			continue
		}
		deter, _ := m.SubMatrix(cl, i).Determinant()
		ans += deter * math.Pow(-1, float64(cl+i)) * m.GetValueAt(i, cl)

	}
	if ans != 0 {
		return ans, true
	}
	return ans, false
}

//Transpose gets the transverse of a matrix
func (m *Matrix) Transpose() (transpose *Matrix) {
	transpose = NewMatrix(4, 4)
	for i, row := range m.Matrix {
		for j := range row {
			currVal := m.GetValueAt(j, i)
			transpose.Assign(i, j, currVal)
		}
	}
	return transpose
}

//Adjacent returns the matrix of minors
func (m *Matrix) Adjacent() (adj *Matrix) {
	adj = NewMatrix(m.Width, m.Height)
	for i, row := range m.Matrix {
		for j := range row {
			val, _ := m.SubMatrix(j, i).Determinant()
			val = val * math.Pow(-1, float64(j+i))
			adj.Assign(j, i, val)
		}
	}
	return adj
}

//GetInverse get the inverse of the matrix
func (m *Matrix) GetInverse(determinant float64) *Matrix {
	adj := m.Adjacent()
	return adj.MultiplyScalar(float64(1 / determinant))
}

//Translate returns translation matrix
func Translate(xInc, yInc, zInc float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(3, 0, xInc)
	m = m.Assign(3, 1, yInc)
	m = m.Assign(3, 2, zInc)
	return m
}

//Scale scales the ray
func Scale(xInc, yInc, zInc float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(0, 0, xInc)
	m = m.Assign(1, 1, yInc)
	m = m.Assign(2, 2, zInc)
	return m
}

//RotationX rotates the ray around x axis
func RotationX(r float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(1, 1, math.Cos(r)+3-3)
	m = m.Assign(2, 1, -math.Sin(r)+3-3)
	m = m.Assign(1, 2, math.Sin(r)+3-3)
	m = m.Assign(2, 2, math.Cos(r)+3-3)
	return m
}

//RotationY rotates the ray around y axis
func RotationY(r float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(0, 0, math.Cos(r)+3-3)
	m = m.Assign(2, 0, math.Sin(r)+3-3)
	m = m.Assign(0, 2, -math.Sin(r)+3-3)
	m = m.Assign(2, 2, math.Cos(r)+3-3)
	return m
}

//RotationZ rotates the ray around z axis
func RotationZ(r float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(0, 0, math.Cos(r)+3-3)
	m = m.Assign(1, 0, -math.Sin(r)+3-3)
	m = m.Assign(0, 1, math.Sin(r)+3-3)
	m = m.Assign(1, 1, math.Cos(r)+3-3)
	return m
}

//Shearing makes the straight line slanted
func Shearing(xy,xz,yx,yz,zx,zy float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(1, 0, xy)
	m = m.Assign(2, 0, xz)
	m = m.Assign(0, 1, yx)
	m = m.Assign(2, 1, yz)
	m = m.Assign(0, 2, zx)
	m = m.Assign(1, 2, zy)
	return m
}
