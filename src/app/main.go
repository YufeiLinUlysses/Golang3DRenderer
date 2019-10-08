package main

import (
	"fmt"
	"reflect"
	//"method"
)

//Test type
type Test struct {
	A int
}

func main() {
	t := &Test{A: 2}
	func do(i interface{}) {
		switch v := i.(type) {
		case *Test:
			fmt.Printf("%v\n",i)
		case int:
			fmt.Printf("Twice %v is %v\n", v, v*2)
		case string:
			fmt.Printf("%q is %v bytes long\n", v, len(v))
		default:
			fmt.Printf("I don't know about type %T!\n", v)
		}
	}
	
	//method.ThirdImage("../../output/test3")

}
