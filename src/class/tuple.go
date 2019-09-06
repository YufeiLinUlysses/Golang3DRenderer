package class

import (
	"fmt"
	"math"
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
func (tone Tuple) Subtract(ttwo Tuple) (ans Tuple, typeOfTuple bool) {
	ans.X = tone.X - ttwo.X
	ans.Y = tone.Y - ttwo.Y
	ans.Z = tone.Z - ttwo.Z
	ans.W = tone.W - ttwo.W
	if ans.W != 1 {
		typeOfTuple = false
	} else {
		typeOfTuple = true
	}
	return ans, typeOfTuple
}

//Mutiply helps the tuple mutiplies a scalar
func (t Tuple) Multiply(scalar float64) Tuple {
	var ans Tuple
	ans.X = t.X * scalar
	ans.Y = t.Y * scalar
	ans.Z = t.Z * scalar
	ans.W = t.W * scalar
	return ans
}

//Divides helps the tuple divide by a scalar
func (t Tuple) Divide(scalar float64) Tuple {
	var ans Tuple
	if scalar != 0 {
		ans.X = t.X / scalar
		ans.Y = t.Y / scalar
		ans.Z = t.Z / scalar
		ans.W = t.W / scalar
	} else {
		fmt.Println("You are dividing by 0")
		return ans
	}
	return ans
}

//Magnitude calculates the magnitude of a vector
func (t Tuple) Magnitude() (mag float64, vecOrNot bool) {
	if t.W != 1 {
		mag = math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z)
		return mag, true
	} else {
		mag = 0
		return mag, false
	}

}

//Normalize normalizes the vector by dividing each element with magnitude
func (t Tuple) Normalize() (ans Tuple, normalized bool) {
	mag, vecOrNot := t.Magnitude()
	if vecOrNot && mag != 0 {
		normalized = true
		ans.X = t.X / mag
		ans.Y = t.Y / mag
		ans.Z = t.Z / mag
		return ans, normalized
	} else if vecOrNot && mag == 0 {
		normalized = false
		fmt.Println("A zero vector")
		return ans, normalized
	} else {
		normalized = false
		return ans, normalized
	}
}

//DotProduct calculates the dot product of two vectors
func (tone Tuple) DotProduct(ttwo Tuple) (ans float64, dotted bool) {
	if tone.W != 1 && ttwo.W != 1 {
		return tone.X*ttwo.X + tone.Y*ttwo.Y + tone.Z * ttwo.Z, true
	} else {
		return 0, false
	}
}

//CrossProduct calculates the cross product of two vectors
func (tone Tuple) CrossProduct(ttwo Tuple) (ans Tuple, crossed bool) {
	if tone.W != 1 && ttwo.W != 1 {
		ans.X = tone.Y*ttwo.Z - tone.Z*ttwo.Y
		ans.Y = tone.Z*ttwo.X - tone.X*ttwo.Z
		ans.Z = tone.X*ttwo.Y - tone.Y*ttwo.X
		return ans, true
	} else {
		return ans, false
	}
}

func main() {
	vone := Tuple{1, 2, 3, 0}
	vtwo := Tuple{2, 3, 4, 0}
	fmt.Println(vone.CrossProduct(vtwo))
}
