package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func ascending(i int) int {
	strNum := strconv.Itoa(i)
	splitNum := strings.Split(strNum, "")
	sort.Sort(sort.Reverse(sort.StringSlice(splitNum)))
	ans, _ := strconv.Atoi(strings.Join(splitNum, ""))
	return ans
}

func descending(i int) int {
	strNum := strconv.Itoa(i)
	splitNum := strings.Split(strNum, "")
	sort.Strings(splitNum)
	ans, _ := strconv.Atoi(strings.Join(splitNum, ""))
	return ans
}

func main() {
	var num int
	//num = 132
	fmt.Scan(&num)

	fmt.Println(ascending(num))
	fmt.Println(descending(num))

}
