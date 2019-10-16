package feature

//Computations type
type Computations struct {
	T         float64
	Shape     interface{}
	Point     Tuple
	OverPoint Tuple
	Eye       Tuple
	Normal    Tuple
	Reflect   Tuple
	Inside    bool
}
