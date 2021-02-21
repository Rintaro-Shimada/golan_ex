package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var num int
	num = 123
	strN := strconv.Itoa(num)
	splitN := strings.Split(strN, "")
	sort.Strings(splitN)

	fmt.Println(splitN)
}
