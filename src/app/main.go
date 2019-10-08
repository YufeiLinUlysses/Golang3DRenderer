package main

import (
	"encoding/json"
	"feature"
	"fmt"
	//"method"
)

func main() {
	//method.ThirdImage("../../output/test3")
	w := feature.DefaultWorld()
	var temp feature.Sphere
	json.Unmarshal(w.Objects[0], &temp)
	fmt.Println(temp)
	// temp := feature.NewTest(2)
	// t, _ := json.Marshal(*temp)
	// var haha feature.Test
	// fmt.Println(temp)
	// fmt.Println(string(t))
	// fmt.Println(json.Unmarshal(t, &haha))
	// fmt.Println(haha)
}
