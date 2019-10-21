package feature

//Computations type
type Computations struct {
	T         float64
	Refract1  float64
	Refract2  float64
	Shape     interface{}
	Point     Tuple
	OverPoint Tuple
	Eye       Tuple
	Normal    Tuple
	Reflect   Tuple
	Inside    bool
}
