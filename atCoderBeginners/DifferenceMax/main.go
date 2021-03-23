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

func check(a,b,c,d int)  {
	var xMax int
	if a<b{
		xMax = b
	}else if  a>b{
		xMax = a
	}else {
		xMax = a
	}

	var yMin int
	if c<d{
		yMin = c
	}else if  c>d{
		yMin = d
	}else {
		yMin = d
	}

	fmt.Println(xMax - yMin)
}

func main() {
	sc.Split(bufio.ScanWords)

	a := nextInt()
	b := nextInt()
	c := nextInt()
	d := nextInt()

	check(a,b,c,d)
}
