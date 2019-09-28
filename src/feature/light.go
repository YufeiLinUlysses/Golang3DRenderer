package feature

//Light type
type Light struct {
	Position  Tuple
	Intensity Color
}

//NewLight establishes a new light instance
func NewLight() *Light {
	l := &Light{
		Position:  *Point(0, 0, 0),
		Intensity: *NewColor(1, 1, 1),
	}
	return l
}

//GetLight gets the information of the light
func (l *Light) GetLight() (np *Tuple, ni *Color) {
	return &l.Position, &l.Intensity
}

//PointLight sets and gets the position and intensity of the light
func (l *Light) PointLight(p Tuple, i Color) Light {
	l.Position = p
	l.Intensity = i
	return *l
}
