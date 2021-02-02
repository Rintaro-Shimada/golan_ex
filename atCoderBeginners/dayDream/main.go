package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func readLine() string {
	sc.Scan()
	return sc.Text()
}

func check(s string) string {
	words := []string{"dream", "dreamer", "erase", "eraser"}

	for len(s) > 0 {
		flag := 0
		for _, w := range words {
			if strings.HasSuffix(s, w) {
				s = s[:len(s)-len(w)]
				break
			}
			flag++
		}
		if flag == len(words) {
			return "NO"
		}
	}
	return "YES"
}

func main() {
	var s string
	s = readLine()
	fmt.Println(check(s))
}
