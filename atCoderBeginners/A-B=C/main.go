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

	for i := 0; i < n; i++ {
		left := nextInt()
		right := nextInt()

		max := (right - (left * 2)) + 1

		if max <= 0 {
			fmt.Println("0")
		} else {
			fmt.Println((max + 1) * max / 2)
		}
	}
}
