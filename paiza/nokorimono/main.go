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

func nextFloat() float64 {
	sc.Scan()
	num, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	return num
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextFloat()
	p := nextFloat()
	q := nextFloat()

	pre := n * (100 - p)/100
	ans := pre*(100 - q)/100

	fmt.Println(ans)
}
