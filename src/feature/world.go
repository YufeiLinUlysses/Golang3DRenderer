package feature

//World type
type World struct {
	Light   []Light
	Spheres []Sphere
}

//NewWorld establishes a new world instance, if nothing is given, it returns nothing
func NewWorld(l []Light, s []Sphere) *World {
	w := &World{
		Light:   l,
		Spheres: s,
	}
	return w
}

//DefaultWorld establishes the default world in the book
func DefaultWorld() *World {
	var lights []Light
	var spheres []Sphere

	light := NewLight()
	lights = append(lights, *light)

	s1 := NewSphere()
	s1.Material.Col = *NewColor(0.8, 1.0, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2

	s2 := NewSphere()
	s2.Transform = Scale(0.5, 0.5, 0.5)

	spheres = append(spheres, *s1, *s2)
	w := NewWorld(lights, spheres)
	return w
}
