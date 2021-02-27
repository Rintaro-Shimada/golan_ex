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

	numA := nextFloat()
	numB := nextFloat()

	fmt.Println((1 - (numB / numA)) * 100)
}
