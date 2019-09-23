package class

//Ray type
type Ray struct {
	Origin    Tuple
	Direction Tuple
}

//NewRay establishes a new Ray instance, overload for inputting two tuples
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

//Translate translates a ray
func (r *Ray) Translate(xInc, yInc, zInc float64) *Ray {
	transV := Vector(xInc, yInc, zInc)
	r.Origin = r.Origin.Add(transV)
	return r
}

//Scale scales the ray
func (r *Ray) Scale(xMult, yMult, zMult float64) *Ray {
	r.Origin = *Point(xMult*r.Origin.X, yMult*r.Origin.Y, zMult*r.Origin.Z)
	r.Direction = *Vector(xMult*r.Direction.X, yMult*r.Direction.Y, zMult*r.Direction.Z)
	return r
}
