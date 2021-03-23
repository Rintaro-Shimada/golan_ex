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

func check(a, b, c, d, x int) int {
	distance := x - a
	if distance == 0 {
		return b + d
	} else if distance < 0 {
		return b
	}

	if distance/c == 0 {
		return b + d
	}
	return b + d + ((distance / c) * d)
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	x := nextInt()

	price := make([]int, n)

	for i := 0; i < n; i++ {
		a := nextInt()
		b := nextInt()
		c := nextInt()
		d := nextInt()
		price[i] = check(a, b, c, d, x)
	}
	sort.Ints(price)
	fmt.Print(price[0])
	fmt.Print(" ")
	fmt.Println(price[n-1])
}
