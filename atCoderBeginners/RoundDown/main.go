package main

import (
	"fmt"
	"strings"
)

func main() {
	var num string

	fmt.Scan(&num)
	splitNum := strings.Split(num, ".")

	fmt.Println(splitNum[0])
}
