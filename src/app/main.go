package main

import (
	"class"
	"fmt"
)

//

func main() {
	m := class.NewLight()
	fmt.Println(m.PointLight(*class.Point(0,0,0),*class.NewColor(1,1,1)))
}
