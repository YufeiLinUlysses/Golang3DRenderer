package main

import (
	"feature"
	"fmt"
	//"method"
)

func main() {
	//method.ThirdImage("../../output/test3")
	w := feature.DefaultWorld()
	this := w.Objects["s1"].(*feature.Sphere)
	fmt.Println(this.Material )
	// temp := feature.NewTest(2)
	// t, _ := json.Marshal(*temp)
	// var haha feature.Test
	// fmt.Println(temp)
	// fmt.Println(string(t))
	// fmt.Println(json.Unmarshal(t, &haha))
	// fmt.Println(haha)
}
