package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readLineFloat64() float64 {
	sc.Scan()
	sc.Text()
	num, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	return num
}

func main() {
	var n float64
	n = readLineFloat64()

	fmt.Println(n)
}
