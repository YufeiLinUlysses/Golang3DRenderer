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

//Translate returns translation matrix
func (r *Ray) Translate(xInc, yInc, zInc float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(3, 0, xInc)
	m = m.Assign(3, 1, yInc)
	m = m.Assign(3, 2, zInc)
	return m
}

//Scale scales the ray
func (r *Ray) Scale(xInc, yInc, zInc float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(0, 0, xInc)
	m = m.Assign(1, 1, yInc)
	m = m.Assign(2, 2, zInc)
	return m
}

//Transform transforms the matrix with a given command
func (r *Ray) Transform(matrix *Matrix) *Ray {
	newR := NewRay(0, 0, 0, 0, 0, 0)
	ori, _ := matrix.MultiplyTuple(&r.Origin)
	newR.Origin = *ori
	dir, _ := matrix.MultiplyTuple(&r.Direction)
	newR.Direction = *dir
	return newR
}
