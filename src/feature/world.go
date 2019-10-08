package feature

import "encoding/json"

//World type
type World struct {
	Light   []Light
	Objects [][]byte
}

//NewWorld establishes a new world instance, if nothing is given, it returns nothing
func NewWorld(l []Light, s [][]byte) *World {
	w := &World{
		Light:   l,
		Objects: s,
	}
	return w
}

//DefaultWorld establishes the default world in the book
func DefaultWorld() *World {
	var lights []Light
	var objects [][]byte

	light := NewLight()
	lights = append(lights, *light)

	s1 := NewSphere()
	s1.Material.Col = *NewColor(0.8, 1.0, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2
	t1, _ := json.Marshal(s1)

	s2 := NewSphere()
	s2.Transform = Scale(0.5, 0.5, 0.5)
	t2, _ := json.Marshal(s2)

	objects = append(objects, t1, t2)
	w := NewWorld(lights, objects)
	return w
}
