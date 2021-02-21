package main

import "fmt"

func check(num int) int {
	var count int

	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			if num < i*j {
				break
			}
			for k := 1; k <= num; k++ {
				pre := i * j * k
				if num >= pre && 0 < pre {
					//fmt.Println(i, j, k)
					count++
				}
			}
		}
	}
	return count
}

func main() {
	var num int
	fmt.Scan(&num)

	fmt.Println(check(num))
}
