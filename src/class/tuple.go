package class

import (
	"fmt"
)

//Tuple Class
type Tuple struct {
	X, Y, Z, W float64
}

//GetTuple tells whether the input of the vector is correct or not
func (v Tuple) GetTuple() (x, y, z float64, typeOfTuple bool) {
	if v.W == 1.0 {
		return v.X, v.Y, v.Z, true
	}
	return v.X, v.Y, v.Z, false
}

//Point checks whether the tuple is a point
func Point(x, y, z float64) Tuple {
	v := Tuple{x, y, z, 1.0}
	return v
}

//Vector checks whether the tuple is a vector
func Vector(x, y, z float64) Tuple {
	v := Tuple{x, y, z, 0}
	return v
}

//Add adds two tuple together
func (tone Tuple) Add(ttwo Tuple) Tuple {
	var ans Tuple
	ans.X = tone.X + ttwo.X
	ans.Y = tone.Y + ttwo.Y
	ans.Z = tone.Z + ttwo.Z
	ans.W = tone.W + ttwo.W
	return ans
}

//Subtract subtracts ttwo from tone
func (tone Tuple) Subtract(ttwo Tuple) Tuple {
	var ans Tuple
	ans.X = tone.X - ttwo.X
	ans.Y = tone.Y - ttwo.Y
	ans.Z = tone.Z - ttwo.Z
	ans.W = tone.W + ttwo.W
	return ans
}

func main() {
	vone := Tuple{4.3, -4.2, 3.1, 1}
	vtwo := Tuple{4.3, -4.2, 3.1, 1}
	fmt.Println(vone.Add(vtwo))
}
