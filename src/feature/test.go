package feature

//Test type
type Test struct {
	a float64
}

//NewTest establishes a new test instance
func NewTest(haha float64) *Test {
	t := &Test{
		a: haha,
	}
	return t
}
