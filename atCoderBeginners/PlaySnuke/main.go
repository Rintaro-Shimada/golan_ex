package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func check(matrix [][]int) int {
	count := 0
	for i := 0; i < len(matrix); i++ {
		pre := matrix[i][0]

		matrix[i][2] = matrix[i][2] - pre

		if matrix[0][2] < 1 {
			return -1
		}

		if matrix[i][2] < 1 {
			count = i
			break
		}
	}

	//fmt.Println(matrix)

	price := make([]int, count)
	for i := 0; i < count; i++ {
		price[i] = matrix[i][1]
	}
	//fmt.Println(price)

	sort.IntSlice(price).Sort()

	//fmt.Println(price)
	return price[0]
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		a := nextInt()
		p := nextInt()
		x := nextInt()
		matrix[i] = []int{a, p, x}
	}
	//fmt.Println(matrix)

	sort.SliceStable(matrix, func(i, j int) bool {
		return matrix[i][0] < matrix[j][0]
	})

	fmt.Println(check(matrix))
}
