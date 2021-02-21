package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

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
	n := 3

	array := make([]int, n)
	for i := 0; i < n; i++ {
		numFloat := nextFloat()
		array[i] = int(math.Round(numFloat * 10000))
	}

	fmt.Println(array)

	xMin := array[0] - array[2]
	xMax := array[0] + array[2]

	yMin := array[1] - array[2]
	yMax := array[1] + array[2]

	fmt.Println(xMin, xMax, yMin, yMax)

	//切り上げ
	xMinInt := xMin / 10000
	xMaxInt := xMax / 10000

	fmt.Println(xMinInt, xMaxInt)
}
