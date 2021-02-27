package main

import (
	"fmt"
	"playground/example_package/mylib"
)

func main() {
	i := []int{1, 2, 3, 4, 5}
	ans := mylib.average(i)
	fmt.Println(ans)
}
