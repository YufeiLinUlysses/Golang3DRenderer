package main

import (
	"feature"
	"fmt"
	//"method"
)

//Test type
type Test struct {
	A int
}

func main() {
	//t := &Test{A: 2}
	w := feature.DefaultWorld()
	fmt.Println(w.Objects["s1"])
	do(w.Objects["s1"])
	//method.ThirdImage("../../output/test3")

}

func do(i interface{}) {
	switch v := i.(type) {
	case *Test:
		fmt.Println(v)
	case *feature.Sphere:
		fmt.Println(v.Material)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
