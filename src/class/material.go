package class

//Material type
type Material struct {
	Col     Color
	Diffuse float64
}

//NewMaterial establishes a new instance for material class
func NewMaterial() *Material {
	m := &Material{
		Col:     *NewColor(1, 1, 1),
		Diffuse: 1,
	}
	return m
}

//GetMaterial gets the information of a material instance
func (m *Material) GetMaterial() (col Color, dif float64) {
	return m.Col, m.Diffuse
}
