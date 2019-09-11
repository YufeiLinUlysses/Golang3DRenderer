package class

//Ray type
type Ray struct {
	Tuple
	Origin    Tuple
	Direction Tuple
}

//NewRay establishes a new Ray instance
func NewRay(oriX, oriY, oriZ, dirX, dirY, dirZ float64) *Ray {
	inputDir, _ := Vector(dirX, dirY, dirZ).Normalize()
	r := &Ray{
		Origin:    *Point(oriX, oriY, oriZ),
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
