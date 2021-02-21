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
	n := nextInt()

	array := make([]float64, n)
	for i := 0; i < n; i++ {
		numFloat := nextFloat()
		array[i] = numFloat
	}

	fmt.Println(array)
}
