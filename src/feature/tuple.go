package feature

import (
	"fmt"
	"math"
)

//Tuple feature
type Tuple struct {
	X, Y, Z, W float64
}

//NewTuple helps to establish a new tuple instance
func NewTuple(x, y, z, w float64) *Tuple {
	t := &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
	return t
}

//Point checks whether the tuple is a point
func Point(x, y, z float64) *Tuple {
	v := NewTuple(x, y, z, 1.0)
	return v
}

//Vector checks whether the tuple is a vector
func Vector(x, y, z float64) *Tuple {
	v := NewTuple(x, y, z, 0)
	return v
}

//GetTuple tells whether the input of the vector is correct or not
func (t *Tuple) GetTuple() (x, y, z float64, typeOfTuple bool) {
	if t.W == 1.0 {
		return t.X, t.Y, t.Z, true
	}
	return t.X, t.Y, t.Z, false
}

//Add adds two tuple together
func (tone *Tuple) Add(ttwo *Tuple) Tuple {
	var ans Tuple
	ans.X = tone.X + ttwo.X
	ans.Y = tone.Y + ttwo.Y
	ans.Z = tone.Z + ttwo.Z
	ans.W = tone.W + ttwo.W
	return ans
}

//Subtract subtracts ttwo from tone
func (tone *Tuple) Subtract(ttwo *Tuple) (ans Tuple, typeOfTuple bool) {
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
func (t *Tuple) Multiply(scalar float64) Tuple {
	var ans Tuple
	ans.X = t.X * scalar
	ans.Y = t.Y * scalar
	ans.Z = t.Z * scalar
	ans.W = t.W * scalar
	return ans
}

//Divides helps the tuple divide by a scalar
func (t *Tuple) Divide(scalar float64) Tuple {
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
func (t *Tuple) Magnitude() (mag float64, vecOrNot bool) {
	if t.W != 1 {
		mag = math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z)
		return mag, true
	}
	mag = 0
	return mag, false
}

//MagnitudeSquared calculates the square of  magnitude of a vector
func (t *Tuple) MagnitudeSquared() (magSq float64, vecOrNot bool) {
	if t.W != 1 {
		magSq = t.X*t.X + t.Y*t.Y + t.Z*t.Z
		return magSq, true
	}
	magSq = 0
	return magSq, false
}

//Normalize normalizes the vector by dividing each element with magnitude
func (t *Tuple) Normalize() (ans Tuple, normalized bool) {
	mag, vecOrNot := t.Magnitude()
	if vecOrNot && mag != 0 {
		normalized = true
		ans.X = t.X / mag
		ans.Y = t.Y / mag
		ans.Z = t.Z / mag
		return ans, normalized
	} else if vecOrNot && mag == 0 {
		normalized = false
		return ans, normalized
	} else {
		normalized = false
		return ans, normalized
	}
}

//DotProduct calculates the dot product of two vectors
func (tone *Tuple) DotProduct(ttwo *Tuple) (ans float64, dotted bool) {
	if tone.W != 1 && ttwo.W != 1 {
		return tone.X*ttwo.X + tone.Y*ttwo.Y + tone.Z*ttwo.Z, true
	}
	return 0, false
}

//CrossProduct calculates the cross product of two vectors
func (tone *Tuple) CrossProduct(ttwo *Tuple) (ans Tuple, crossed bool) {
	if tone.W != 1 && ttwo.W != 1 {
		ans.X = tone.Y*ttwo.Z - tone.Z*ttwo.Y
		ans.Y = tone.Z*ttwo.X - tone.X*ttwo.Z
		ans.Z = tone.X*ttwo.Y - tone.Y*ttwo.X
		return ans, true
	} else {
		return ans, false
	}
}