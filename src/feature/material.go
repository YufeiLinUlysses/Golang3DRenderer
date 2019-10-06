package feature

//Material type
type Material struct {
	Col     Color
	Ambient float64
	Diffuse float64
	Specular float64
	Shininess float64
}

//NewMaterial establishes a new instance for material feature
func NewMaterial() *Material {
	m := &Material{
		Col:     *NewColor(1, 1, 1),
		Ambient: 0.1,
		Diffuse: 0.9,
		Specular: 0.9,
		Shininess: 200,
	}
	return m
}

//GetMaterial gets the information of a material instance
func (m *Material) GetMaterial() (col Color, dif float64) {
	return m.Col, m.Diffuse
}

//Lighting gets the lighting of the object and decides the color of a pixel
func (m *Material) Lighting(l Light, p, n *Tuple)(col Color){
	dir,_ := l.Position.Subtract(p)
	ndir,_ := dir.Normalize()
	dP,_ := ndir.DotProduct(n)
	if dP >= 0{
		m.Col = m.Col.Multiply(dP)
	}else{
		m.Col = *NewColor(0,0,0)
	}
	return m.Col
}
