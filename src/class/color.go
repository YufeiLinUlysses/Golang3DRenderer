package class

type Color struct {
	Tuple
	R, G, B float64
}

func NewColor(red, green, blue float64) *Color {
	c := &Color{
		Tuple: Tuple{X: red, Y: green, Z: blue, W: 0},
		R:     red,
		G:     green,
		B:     blue,
	}
	return c
}

func (c *Color) GetColor() (r, g, b float64) {
	r, g, b, _ = c.GetTuple()
	return r, g, b
}

func (cone *Color) Add(ctwo *Color) Color {
	ans := cone.Tuple.Add(&ctwo.Tuple)
	fin := *NewColor(ans.X, ans.Y, ans.Z)
	return fin
}


