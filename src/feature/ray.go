package feature

/*Ray type contains all necessary component of a ray
 *Ray contains tuple*/
type Ray struct {
	Origin    Tuple
	Direction Tuple
}

/*NewRay establishes a new Ray instance, overload for inputting two tuples
 *NewRay takes in two tuple
 *NewRay returns a ray*/
func NewRay(origin Tuple, direction Tuple) *Ray {
	inputDir, _ := direction.Normalize()
	r := &Ray{
		Origin:    origin,
		Direction: inputDir,
	}
	return r
}

/*GetRay gets the orign and direction of a ray
 *GetRay can only be called by a ray
 *GetRay returns two tuple*/
func (ray *Ray) GetRay() (ori, dir Tuple) {
	return ray.Origin, ray.Direction
}

/*Position computes a point from a distance
 *Position can only be called by a ray
 *Position takes in a float
 *Poistion returns a tuple*/
func (ray *Ray) Position(dist float64) Tuple {
	var ans Tuple
	ans.X = ray.Origin.X + dist*ray.Direction.X
	ans.Y = ray.Origin.Y + dist*ray.Direction.Y
	ans.Z = ray.Origin.Z + dist*ray.Direction.Z
	ans.W = 1
	return ans
}

/*Transform transforms the matrix with a given command
 *Transform can only be called by a ray
 *Transform takes in a matrix
 *Transform returns a ray*/
func (ray *Ray) Transform(matrix *Matrix) *Ray {
	newR := NewRay(*Point(0, 0, 0), *Vector(0, 0, 0))
	ori, _ := matrix.MultiplyTuple(&ray.Origin)
	newR.Origin = *ori
	dir, _ := matrix.MultiplyTuple(&ray.Direction)
	newR.Direction = *dir
	return newR
}
