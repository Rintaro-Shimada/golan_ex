package main

import "fmt"

func check(num int) int {
	var count int

	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			if i*j > num {
				break
			}
			for k := 1; k <= num; k++ {
				if i*j*k <= num {
					count++
				} else {
					break
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
