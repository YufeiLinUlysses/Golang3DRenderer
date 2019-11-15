package main

import (
	"feature"
	"fmt"
	//"method"
)

func main() {
	s := feature.NewSphere()
	c := feature.NewCube()
	csg := feature.NewCSG("union",s,c)
	fmt.Println(s.Parent)
	fmt.Println(c.Parent)
	fmt.Println(csg.Left)
}
