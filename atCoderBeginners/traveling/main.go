package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	sc.Split(bufio.ScanWords)

	var n int
	n = nextInt()
	var t, x, y int
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		t = nextInt()
		x = nextInt()
		y = nextInt()
		matrix[i] = []int{t, x, y}
	}

	//位置情報
	position := []int{0, 0}

	var status = false
	for i := 0; i < n; i++ {
		//移動回数
		workCount := matrix[i][0] //3
		if i > 0 {
			workCount = matrix[i][0] - matrix[i-1][0]
		}

		for j := 0; j < workCount; j++ {
			if position[0] < matrix[i][1] {
				position[0]++
			} else if position[1] < matrix[i][2] {
				position[1]++
			} else if position[0] > matrix[i][1] {
				position[0]--
			} else if position[1] > matrix[i][2] {
				position[1]--
			} else {
				position[0]++
			}
			//fmt.Println(position)
		}

		if position[0] == matrix[i][1] && position[1] == matrix[i][2] {
			status = true
			continue
		} else {
			status = false
			break
		}
	}

	if status {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
