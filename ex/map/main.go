package main

import "fmt"

func main() {
	max := 4
	matrix := make([][]int, max)

	for i := 0; i < max; i++ {
		matrix[i] = []int{i, i, i}
	}

	fmt.Println(matrix)
}
