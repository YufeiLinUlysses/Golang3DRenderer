package main

import (
	"class"
	"fmt"
	//"fmt"
	//"method"
)

func main() {
	//method.SecondImage("../../output/test2")
	tables := []struct {
		w1, h1   int
		value1   []float64
		w2, h2   int
		value2   []float64
	}{
		{4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 0, 0, 0, 1}, 4, 4, []float64{-2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 0, 0, 0, 1}},
	}
	
	for _, table := range tables {
		count := 0
		m := class.NewMatrix(table.w1, table.h1)
		for i, row := range m.Matrix {
			for j := range row {
				m = m.Assign(j, i, table.value1[count])
				count++
			}
		}
		count = 0
		m2 := class.NewMatrix(table.w2, table.h2)
		for i, row := range m2.Matrix {
			for j := range row {
				m2 = m2.Assign(j, i, table.value2[count])
				count++
			}
		}
		ansM,_ := m.Multiply(m2)
		fmt.Println(ansM)
	}
}
