package feature

//Ray type
type Ray struct {
	Origin    Tuple
	Direction Tuple
}

//NewRay establishes a new Ray instance, overload for inputting two tuples
func NewRay(origin Tuple, direction Tuple) *Ray {
	inputDir, _ := direction.Normalize()
	r := &Ray{
		Origin:    origin,
		Direction: inputDir,
	}
	return r
}

//GetRay gets the orign and direction of a ray
func (r *Ray) GetRay() (ori, dir Tuple) {
	return r.Origin, r.Direction
}

//Position computes a point from a distance
func (r *Ray) Position(dist float64) Tuple {
	var ans Tuple
	ans.X = r.Origin.X + dist*r.Direction.X
	ans.Y = r.Origin.Y + dist*r.Direction.Y
	ans.Z = r.Origin.Z + dist*r.Direction.Z
	ans.W = 1
	return ans
}

//Transform transforms the matrix with a given command
func (r *Ray) Transform(matrix *Matrix) *Ray {
	newR := NewRay(*Point(0, 0, 0), *Vector(0, 0, 0))
	ori, _ := matrix.MultiplyTuple(&r.Origin)
	newR.Origin = *ori
	dir, _ := matrix.MultiplyTuple(&r.Direction)
	newR.Direction = *dir
	return newR
}
