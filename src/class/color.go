package class

//Color type
type Color struct {
	Tuple
	R, G, B float64
}

//NewColor establishes a new color instance
func NewColor(red, green, blue float64) *Color {
	c := &Color{
		Tuple: Tuple{X: red, Y: green, Z: blue, W: 0},
		R:     red,
		G:     green,
		B:     blue,
	}
	return c
}

//GetColor helps to get the color vector elements
func (cone *Color) GetColor() (r, g, b float64) {
	r, g, b, _ = cone.GetTuple()
	return r, g, b
}

//Add helps to add two color vectors
func (cone *Color) Add(ctwo *Color) Color {
	ans := cone.Tuple.Add(&ctwo.Tuple)
	fin := *NewColor(ans.X, ans.Y, ans.Z)
	return fin
}

//Subtract helps to subtract two color vectors
func (cone *Color) Subtract(ctwo *Color) Color {
	ans,_ := cone.Tuple.Subtract(&ctwo.Tuple)
	fin := *NewColor(ans.X, ans.Y, ans.Z)
	return fin
} 

//Multiply calculates the vector of a vector multiplies a scalar
func (cone *Color) Multiply(scalar float64) Color {
	ans := cone.Tuple.Multiply(scalar)
	fin := *NewColor(ans.X, ans.Y, ans.Z)
	return fin
}

//ColorMultiply multiplies color from two color vectors
func (cone *Color) ColorMultiply(ctwo *Color) Color{
	ans := *NewColor(cone.R * ctwo.R, cone.G * ctwo.G, cone.B * ctwo.B)
	return ans
}
