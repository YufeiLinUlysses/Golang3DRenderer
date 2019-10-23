package feature

/*Color type contains necessary components of a color
 *Color Inherits from Tuple class*/
type Color struct {
	Tuple
	R, G, B float64
}

/*NewColor establishes a new color instance
 *NewColor takes in three float representing the amount of red, green and blue
 *NewColor returns a color instance*/
func NewColor(red, green, blue float64) *Color {
	col := &Color{
		Tuple: Tuple{X: red, Y: green, Z: blue, W: 0},
		R:     red,
		G:     green,
		B:     blue,
	}
	return col
}

/*GetColor helps to get the color vector elements
 *GetColor can only be called by a color instance
 *GetColor returns the amount of red, green, and blue*/
func (col *Color) GetColor() (r, g, b float64) {
	r, g, b, _ = col.GetTuple()
	return r, g, b
}

/*Add helps to add two color vectors
 *Add could only be called by a color instance
 *Add takes in a second color instance
 *Add returns a color instance*/
func (col *Color) Add(c2 *Color) Color {
	ans := col.Tuple.Add(&c2.Tuple)
	fin := *NewColor(ans.X, ans.Y, ans.Z)
	return fin
}

/*Subtract helps to subtract two color vectors
 *Subtract could only be called by a color instance
 *Subtract takes in a second color instance
 *Subtract returns a color instance*/
func (col *Color) Subtract(c2 *Color) Color {
	ans, _ := col.Tuple.Subtract(&c2.Tuple)
	fin := *NewColor(ans.X, ans.Y, ans.Z)
	return fin
}

/*Multiply calculates the vector of a vector multiplies a scalar
 *Multiply could only be called by a color instance
 *Multiply takes in a float
 *Multiply returns a color instance*/
func (col *Color) Multiply(scalar float64) Color {
	ans := col.Tuple.Multiply(scalar)
	fin := *NewColor(ans.X, ans.Y, ans.Z)
	return fin
}

/*ColorMultiply multiplies color from two color vectors
 *ColorMultiply could only be called by a color instance
 *ColorMultiply takes in a second color instance
 *ColorMultiply returns a color instance*/
func (col *Color) ColorMultiply(c2 *Color) Color {
	ans := *NewColor(col.R*c2.R, col.G*c2.G, col.B*c2.B)
	return ans
}
