package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func readLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	var s string
	s = readLine()

	fmt.Println(s)
}
