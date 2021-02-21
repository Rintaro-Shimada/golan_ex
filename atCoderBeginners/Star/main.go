package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)

	oneH := 100 * (1 + (num / 100))
	ans := oneH - num

	fmt.Println(ans)
}
