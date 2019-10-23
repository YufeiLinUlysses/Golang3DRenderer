package feature

import (
	"math"
)

/*Tuple feature contains all necessary component of a tuple
 *In this implementation we see both a point and a vector as a tuple*/
type Tuple struct {
	X, Y, Z, W float64
}

/*NewTuple helps to establish a new tuple instance
 *NewTuple takes in four float
 *NewTuple returns a tuple*/
func NewTuple(x, y, z, w float64) *Tuple {
	tuple := &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
	return tuple
}

/*Point checks whether the tuple is a point
 *Point is a public function
 *Point takes in three float
 *Point returns a tuple*/
func Point(x, y, z float64) *Tuple {
	v := NewTuple(x, y, z, 1.0)
	return v
}

/*Vector checks whether the tuple is a vector
 *Vector is a public function
 *Vector takes in three float
 *Vector returns a tuple*/
func Vector(x, y, z float64) *Tuple {
	v := NewTuple(x, y, z, 0)
	return v
}

/*GetTuple tells whether the input of the vector is correct or not
 *GetTuple can only be called by a tuple
 *GetTuple returns three float and a bool*/
func (tuple *Tuple) GetTuple() (x, y, z float64, typeOfTuple bool) {
	if tuple.W == 1.0 {
		return tuple.X, tuple.Y, tuple.Z, true
	}
	return tuple.X, tuple.Y, tuple.Z, false
}

/*Add adds two tuple together
 *Add can only be called by a tuple
 *Add takes in a tuple
 *Add returns a tuple*/
func (tuple *Tuple) Add(t2 *Tuple) Tuple {
	var ans Tuple
	ans.X = tuple.X + t2.X
	ans.Y = tuple.Y + t2.Y
	ans.Z = tuple.Z + t2.Z
	ans.W = tuple.W + t2.W
	return ans
}

/*Subtract subtracts ttwo from tone
 *Subtract can only be called by a tuple
 *Subtract takes in a tuple
 *Subtract returns a tuple*/
func (tuple *Tuple) Subtract(t2 *Tuple) (ans Tuple, typeOfTuple bool) {
	ans.X = tuple.X - t2.X
	ans.Y = tuple.Y - t2.Y
	ans.Z = tuple.Z - t2.Z
	ans.W = tuple.W - t2.W
	if ans.W != 1 {
		typeOfTuple = false
	} else {
		typeOfTuple = true
	}
	return ans, typeOfTuple
}

/*Multiply helps the tuple mutiplies a scalar
 *Multiply can only be called by a tuple
 *Multiply takes in a tuple
 *Multiply returns a tuple*/
func (tuple *Tuple) Multiply(scalar float64) Tuple {
	var ans Tuple
	ans.X = tuple.X * scalar
	ans.Y = tuple.Y * scalar
	ans.Z = tuple.Z * scalar
	ans.W = tuple.W * scalar
	return ans
}

/*Divide helps the tuple divide by a scalar
 *Divide can only be called by a tuple
 *Divide takes in a tuple
 *Divide returns a tuple*/
func (tuple *Tuple) Divide(scalar float64) (Tuple, bool) {
	var ans Tuple
	if scalar != 0 {
		ans.X = tuple.X / scalar
		ans.Y = tuple.Y / scalar
		ans.Z = tuple.Z / scalar
		ans.W = tuple.W / scalar
	} else {
		return ans, false
	}
	return ans, true
}

/*Magnitude calculates the magnitude of a vector
 *Magnitude can only be called by a tuple
 *Magnitude returns a float, a bool*/
func (tuple *Tuple) Magnitude() (mag float64, vecOrNot bool) {
	if tuple.W != 1 {
		mag = math.Sqrt(tuple.X*tuple.X + tuple.Y*tuple.Y + tuple.Z*tuple.Z)
		return mag, true
	}
	mag = 0
	return mag, false
}

/*MagnitudeSquared calculates the square of  magnitude of a vector
 *MagnitudeSquared can only be called by a tuple
 *MagnitudeSquared returns a float, a bool*/
func (tuple *Tuple) MagnitudeSquared() (magSq float64, vecOrNot bool) {
	if tuple.W != 1 {
		magSq = tuple.X*tuple.X + tuple.Y*tuple.Y + tuple.Z*tuple.Z
		return magSq, true
	}
	magSq = 0
	return magSq, false
}

/*Normalize normalizes the vector by dividing each element with magnitude
 *Normalize can only be called by a tuple
 *Normalize returns a tuple, a bool*/
func (tuple *Tuple) Normalize() (ans Tuple, normalized bool) {
	mag, vecOrNot := tuple.Magnitude()
	if vecOrNot && mag != 0 {
		normalized = true
		ans.X = tuple.X / mag
		ans.Y = tuple.Y / mag
		ans.Z = tuple.Z / mag
		return ans, normalized
	} else if vecOrNot && mag == 0 {
		normalized = false
		return ans, normalized
	} else {
		normalized = false
		return ans, normalized
	}
}

/*DotProduct calculates the dot product of two vectors
 *DotProduct can only be called by a tuple
 *DotProduct takes in a tuple
 *DotProduct returns a float, a bool*/
func (tuple *Tuple) DotProduct(t2 *Tuple) (ans float64, dotted bool) {
	if tuple.W != 1 && t2.W != 1 {
		return tuple.X*t2.X + tuple.Y*t2.Y + tuple.Z*t2.Z, true
	}
	return 0, false
}

/*CrossProduct calculates the cross product of two vectors
 *CrossProduct can only be called by a tuple
 *CrossProduct takes in a tuple
 *CrossProduct returns a tuple, a bool*/
func (tuple *Tuple) CrossProduct(t2 *Tuple) (ans Tuple, crossed bool) {
	if tuple.W != 1 && t2.W != 1 {
		ans.X = tuple.Y*t2.Z - tuple.Z*t2.Y
		ans.Y = tuple.Z*t2.X - tuple.X*t2.Z
		ans.Z = tuple.X*t2.Y - tuple.Y*t2.X
		return ans, true
	}
	return ans, false
}

/*Reflect takes in an incident vector and finds out the reflected vector
 *Reflect can only be called by a tuple
 *Reflect takes in a tuple
 *Reflect returns a tuple, a bool*/
func (tuple *Tuple) Reflect(normal *Tuple) (ans Tuple, reflected bool) {
	dotted, _ := tuple.DotProduct(normal)
	processed := normal.Multiply(dotted * 2)
	return tuple.Subtract(&processed)
}
