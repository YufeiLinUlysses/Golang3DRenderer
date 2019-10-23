package feature

/*Light type contains all necessary components of a light instance
 *Light contains a tuple and a color*/
type Light struct {
	Position  Tuple
	Intensity Color
}

/*NewLight establishes a new light instance
 *NewLight returns a light that is at the origin and the color white*/
func NewLight() *Light {
	lig := &Light{
		Position:  *Point(0, 0, 0),
		Intensity: *NewColor(1, 1, 1),
	}
	return lig
}

/*GetLight gets the information of the light
 *GetLight can only be called by a light instance
 *GetLight returns a tuple and a color*/
func (lig *Light) GetLight() (np *Tuple, ni *Color) {
	return &lig.Position, &lig.Intensity
}

/*PointLight sets and gets the position and intensity of the light
 *PointLight can only be called by a light
 *PointLight takes in a tuple and a color*/
func (lig *Light) PointLight(pos Tuple, intens Color) Light {
	lig.Position = pos
	lig.Intensity = intens
	return *lig
}
