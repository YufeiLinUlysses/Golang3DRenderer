package main

import (
	"class"
	"fmt"
	//"fmt"
	//"method"
)

func main() {
	//method.SecondImage("../../output/test2")
	// tables := []struct {
	// 	w1, h1 int
	// 	value1 []float64
	// 	w2, h2 int
	// 	value2 []float64
	// }{
	// 	{4, 4, []float64{1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1}, 4, 4, []float64{-2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 0, 0, 0, 1}},
	// }

	// for _, table := range tables {
	// 	count := 0
	// 	m := class.NewMatrix(table.w1, table.h1)
	// 	for i, row := range m.Matrix {
	// 		for j := range row {
	// 			m = m.Assign(j, i, table.value1[count])
	// 			count++
	// 		}
	// 	}
	// 	//identity, _ := m.GetIdentity()
	// 	// m2 := class.NewMatrix(table.w2, table.h2)
	// 	// for i, row := range m2.Matrix {
	// 	// 	for j := range row {
	// 	// 		m2 = m2.Assign(j, i, table.value2[count])
	// 	// 		count++
	// 	// 	}
	// 	// }
	// 	ansM, _ := m.MultiplyTuple(class.NewTuple(1, 2, 3, 1))
	// 	fmt.Println(ansM)
	// }
	count := 0
	value := []float64{1, 2, 6, -5, 8, -4, 2, 6, 4}
	m := class.NewMatrix(3, 3)
	for i, row := range m.Matrix {
		for j := range row {
			m = m.Assign(j, i, value[count])
			count++
		}
	}
	fmt.Println(m.Determinant())
}
