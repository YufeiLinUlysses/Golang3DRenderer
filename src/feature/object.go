package feature

//Object type
type Object struct {
	Addr []byte
}

//SetAddr sets the addr
func (o *Object) SetAddr(json []byte) *Object {
	o.Addr = json
	return o
}

//Object Class
//Scale
//Translate
//Transform
//Material
//Virtual IntersectionWidth
