package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	for i := 0; i < len(s)+2; i++ {
		fmt.Print("+")
	}
	fmt.Print("\n")
	fmt.Print("+")
	fmt.Print(s)
	fmt.Print("+")
	fmt.Print("\n")
	for i := 0; i < len(s)+2; i++ {
		fmt.Print("+")
	}
}
