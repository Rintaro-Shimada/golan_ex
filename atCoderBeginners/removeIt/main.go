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

	n := nextInt()
	x := nextInt()

	nums := make([]int, 0, 0)

	for i := 0; i < n; i++ {
		a := nextInt()
		if x != a {
			nums = append(nums, a)
		}
	}

	if len(nums) != 0 {
		for i := 0; i < len(nums); i++ {
			fmt.Print(nums[i])
			fmt.Print(" ")
		}
	} else {
		fmt.Println("")
	}
}
