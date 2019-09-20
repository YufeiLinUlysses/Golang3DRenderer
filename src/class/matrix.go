package class

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
func (m *Matrix) EqualTo(m2 Matrix) bool{
	if m.Width == m2.Width && m.Height==m2.Height{
		for i, row := range m.Matrix {
			for j := range row {
				mVal := m.GetValueAt(j, i)
				m2Val := m2.GetValueAt(j,i)
				if m2Val != mVal {
					return false
				}
			}
		}
		return true
	}
	return false
}
