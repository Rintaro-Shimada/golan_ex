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

func check(in string) string {
	words := []string{"dream", "dreamer", "erase", "eraser"}

	for len(in) > 0 {
		found := false
		for _, s := range words {
			if !strings.HasSuffix(in, s) {
				continue
			}
			in = in[:len(in)-len(s)]
			fmt.Println("======")
			fmt.Println(in)
			found = true
		}
		if !found {
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
