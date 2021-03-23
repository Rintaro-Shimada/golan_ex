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
	a := nextInt()
	b := nextInt()
	n := nextInt()

	for i := 0; i < n; i++ {
		ans := check(a, b)
		fmt.Println(ans)
	}
}

func check(a, b int) string {
	A := nextInt()
	B := nextInt()

	if a < A {
		return "Low"
	} else if a > A {
		return "High"
	} else if a == A{
		if b < B {
			return "High"
		} else if b > B{
			return "Low"
		}
	}
	return "err"
}
