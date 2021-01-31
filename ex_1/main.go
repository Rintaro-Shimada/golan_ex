package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)
	fmt.Scan(&y)
	ThreePoint(x, y)
}

func ThreePoint(x, y int) {
	var numLow, numHigh int
	if x < y {
		numLow = x
		numHigh = y
	} else {
		numLow = y
		numHigh = x
	}

	if (numLow + 3) > numHigh {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
