package class

import (
	"fmt"
)

//Vertex Class
type Vertex struct {
	X, Y, Z, W float64
}

//GetTuple tells
func (v Vertex) GetTuple() (x, y, z float64, typeOfTuple bool) {
	if v.W == 1.0 {
		return v.X, v.Y, v.Z, true
	}
	return v.X, v.Y, v.Z, false
}

func main() {
	v := Vertex{4.3, -4.2, 3.1, 1.0}
	fmt.Println(v.GetTuple())
}
