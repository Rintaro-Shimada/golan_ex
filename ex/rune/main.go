package main

import "fmt"

func main() {
	s := "abcde"

	for i := 0; i < len(s); i++ {
		b := s[i]      // byte
		fmt.Println(b) // 227, 129, 130...
	}
	fmt.Println("==========================")

	c := "#"[0] // byte
	fmt.Println(c)
	fmt.Println("==========================")

	for _, r := range s {
		// rune
		fmt.Println(r) // 12354, 12356, 12358
	}

	fmt.Println("==========================")

	s2 := "あいう"
	fmt.Println([]rune(s2)) // [12354 12356 12358]
	fmt.Println([]byte(s2)) // [227 129 130 227 129 132 227 129 134]

	s3 := "#"
	fmt.Println([]rune(s3))
	fmt.Println([]byte(s3))

}
