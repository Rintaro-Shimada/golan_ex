package main

import (
	"fmt"
	"strconv"
)

func check(num int) int{
	strNum := strconv.Itoa(num)

	//fmt.Println(len(strNum))
	splitNum := len(strNum) / 2
	//fmt.Println(splitNum)

	//fmt.Println(string([]rune(strNum)[:splitNum]))
	LeftNum := string([]rune(strNum)[:splitNum])

	//fmt.Println(string([]rune(strNum)[splitNum:]))
	rightNum := string([]rune(strNum)[splitNum:])

	leftInt, _ := strconv.Atoi(LeftNum)
	rightInt, _ := strconv.Atoi(rightNum)

	//fmt.Println(leftInt)
	//fmt.Println(rightInt)

	if rightInt == 0{
		return leftInt -1
	}

	if leftInt > rightInt {
		return rightInt
	} else if leftInt < rightInt {
		return leftInt
	}

	return leftInt
}

func main() {
	var num int
	fmt.Scan(&num)

	fmt.Println(check(num))
}
