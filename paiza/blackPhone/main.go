package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)

	var numDistance = []int{12, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	arr := strings.Split(str, "-")
	str = strings.Join(arr, "")
	nums := strings.Split(str, "")
	//num, _ := strconv.Atoi(str)
	//fmt.Println(nums)

	var ans int
	for i:=0;i<10; i++{
		a ,_ := strconv.Atoi(nums[i])
		//fmt.Println(numDistance[a])
		ans += numDistance[a]
	}

	fmt.Println(ans*2)
}
