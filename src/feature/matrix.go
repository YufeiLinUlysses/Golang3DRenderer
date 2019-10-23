package feature

import (
	"math"
)

/*Matrix type contains all necessary components of a matrix*/
type Matrix struct {
	Width, Height int
	Matrix        [][]float64
	determinant   float64
	inverse       [][]float64
	hasInv        bool
	hasDeter      bool
}

/*NewMatrix sets up a new matrix instance
 *NewMatrix takes in the width and height of a matrix
 *NewMatrix returns a matrix*/
func NewMatrix(width, height int) *Matrix {
	mat := make([][]float64, height)
	for i := 0; i < height; i++ {
		mat[i] = make([]float64, width)
	}

	matrix := &Matrix{
		Width:  width,
		Height: height,
		Matrix: mat,
		hasInv: false,
	}
	return matrix
}

/*Assign assigns value for a specific element in a matrix
 *Assign takes in two int and a float
 *Assign returns a matrix*/
func (matrix *Matrix) Assign(cl, rw int, value float64) *Matrix {
	matrix.Matrix[rw][cl] = value
	return matrix
}

/*GetValueAt returns the value of a specific place
 *GetValueAt can only be called by a matrix
 *GetValueAt takes in two int
 *GetValueAt returns a float*/
func (matrix *Matrix) GetValueAt(cl, rw int) float64 {
	return matrix.Matrix[rw][cl]
}

/*EqualTo determines whether two matrix are the same
 *EqualTo could only be called by a matrix
 *EqualTo takes in another matrix
 *EqualTo returns a bool*/
func (matrix *Matrix) EqualTo(m2 Matrix) bool {
	if matrix.Width == m2.Width && matrix.Height == m2.Height {
		for i, row := range matrix.Matrix {
			for j := range row {
				mVal := matrix.GetValueAt(j, i)
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

/*GetIdentity gives the identity matrix of the size of the matrix
 *GetIdentity could be only called by a matrix instance
 *GetIdentity returns a matrix and a bool*/
func (matrix *Matrix) GetIdentity() (identity *Matrix, square bool) {
	if matrix.Width != matrix.Height {
		return identity, false
	}
	identity = NewMatrix(matrix.Width, matrix.Height)
	for i := 0; i < matrix.Height; i++ {
		identity.Assign(i, i, 1)
	}
	return identity, true
}

/*Multiply multiplies two matrices
 *Multiply could only be called by a matrix
 *Multiply takes in a matrix
 *Multiply returns a matrix and a bool*/
func (matrix *Matrix) Multiply(m2 *Matrix) (ansM *Matrix, multiplied bool) {
	ansM = NewMatrix(matrix.Height, m2.Height)
	if matrix.Width != m2.Height {
		return ansM, false
	}
	for i := range matrix.Matrix {
		for j := range m2.Matrix {
			Value := float64(0)
			for k := range ansM.Matrix {
				Value += matrix.GetValueAt(k, i) * m2.GetValueAt(j, k)
			}
			ansM.Assign(j, i, Value)
		}
	}
	return ansM, true
}

/*MultiplyTuple multiplies tuple
 *MultiplyTuple can only be called by a matrix
 *MultiplyTuple takes in a tuple
 *MultiplyTuple returns a tuple and a bool*/
func (matrix *Matrix) MultiplyTuple(tupl *Tuple) (ansT *Tuple, istuple bool) {
	tup := []float64{tupl.X, tupl.Y, tupl.Z, tupl.W}
	var ans []float64
	if matrix.Height != 4 {
		return ansT, false
	}
	for i := 0; i < 4; i++ {
		temp := float64(0)
		for j := 0; j < 4; j++ {
			temp += matrix.GetValueAt(j, i) * tup[j]
		}
		ans = append(ans, temp)
	}
	ansT = NewTuple(ans[0], ans[1], ans[2], ans[3])
	return ansT, true
}

/*MultiplyScalar mutlipliese the matrix to a scalar
 *MultiplyScalar can only be called by a matrix
 *MultiplyScalar takes in a float
 *MultiplyScalar returns a matrix*/
func (matrix *Matrix) MultiplyScalar(scalar float64) *Matrix {
	for i, row := range matrix.Matrix {
		for j := range row {
			currValue := matrix.GetValueAt(j, i)
			matrix.Assign(j, i, scalar*currValue)
		}
	}
	return matrix
}

/*SubMatrix gets the cofactor of a matrix at a certain location
 *SubMatrix can only be called by a matrix
 *SubMatrix takes in two int
 *SubMatrix returns a matrix*/
func (matrix *Matrix) SubMatrix(cl, rw int) *Matrix {
	var value []float64
	ans := NewMatrix(matrix.Width-1, matrix.Height-1)
	for i := range matrix.Matrix {
		if i != cl {
			for j := range matrix.Matrix {
				if j != rw {
					value = append(value, matrix.Matrix[i][j])
				}
			}
		}
	}
	count := 0
	for i := 0; i < matrix.Width-1; i++ {
		for j := 0; j < matrix.Width-1; j++ {
			ans.Assign(j, i, value[count])
			count++
		}
	}
	return ans
}

/*Determinant gets the determinant of a 2x2 square matrix
 *Determinant can only be called by a matrix
 *Determinant returns a float and a bool*/
func (matrix *Matrix) Determinant() (ans float64, invertible bool) {
	var cl int
	if matrix.hasDeter {
		return matrix.determinant, true
	}
	if matrix.Width == 2 {
		ans = matrix.Matrix[0][0]*matrix.Matrix[1][1] - matrix.Matrix[0][1]*matrix.Matrix[1][0]
		return ans, true
	}
	if matrix.Width == 4 {
		cl = 3
	}
	for i := 0; i < len(matrix.Matrix); i++ {
		if matrix.GetValueAt(i, cl) == 0 {
			continue
		}
		deter, _ := matrix.SubMatrix(cl, i).Determinant()
		ans += deter * math.Pow(-1, float64(cl+i)) * matrix.GetValueAt(i, cl)

	}
	if ans != 0 {
		matrix.determinant = ans
		matrix.hasDeter = true
		return ans, true
	}
	return ans, false
}

/*Transpose gets the transverse of a matrix
 *Transpose can only be called by a matrix
 *Transpose returns a matrix*/
func (matrix *Matrix) Transpose() (transpose *Matrix) {
	transpose = NewMatrix(4, 4)
	for i, row := range matrix.Matrix {
		for j := range row {
			currVal := matrix.GetValueAt(j, i)
			transpose.Assign(i, j, currVal)
		}
	}
	return transpose
}

/*Adjacent returns the matrix of minors
 *Adjacent can only be called by a matrix
 *Adjacent returns a matrix*/
func (matrix *Matrix) Adjacent() (adj *Matrix) {
	adj = NewMatrix(matrix.Width, matrix.Height)
	for i, row := range matrix.Matrix {
		for j := range row {
			val, _ := matrix.SubMatrix(j, i).Determinant()
			val = val * math.Pow(-1, float64(j+i))
			adj.Assign(j, i, val)
		}
	}
	return adj
}

/*GetInverse get the inverse of the matrix
 *GetInverse can only be called by a matrix
 *GetInverse takes in a float
 *GetInverse returns a matrix*/
func (matrix *Matrix) GetInverse(determinant float64) *Matrix {
	if matrix.hasInv {
		ans := NewMatrix(4, 4)
		for i, row := range ans.Matrix {
			for j := range row {
				ans.Assign(j, i, matrix.inverse[i][j])
			}
		}
		return ans
	}
	adj := matrix.Adjacent()
	ans := adj.MultiplyScalar(float64(1 / determinant))
	matrix.hasInv = true
	matrix.inverse = ans.Matrix
	return ans
}
