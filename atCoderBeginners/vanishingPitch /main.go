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

	v := nextInt()
	t := nextInt()
	s := nextInt()
	d := nextInt()

	vt := v * t
	vs := v * s

	if d < vt || d > vs {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
