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
		w, h     int
		value    []float64
		ansValue []float64
	}{
		{4, 4, []float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 0, 0, 0, 1}, []float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 0, 0, 0, 1}},
	}
	for _, table := range tables {
		count := 0
		m := class.NewMatrix(table.w, table.h)
		for i, row := range m.Matrix {
			for j := range row {
				m = m.Assign(j, i, table.value[count])
				count++
			}
		}
		fmt.Println(m)
	}
}
