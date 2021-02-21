package main

import (
	"bufio"
	"fmt"
	"os"
)

var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

func readLine() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func check(s string) string {
	for i := 0; i < len(s); i++ {
		if 0 == i%2 {
			if s[i] >= 97 {
				//fmt.Println("奇数 小文字")
			} else {
				return "No"
			}
		}

		if 1 == i%2 {
			if s[i] < 97 {
				//fmt.Println("偶数 大文字")
			} else {
				return "No"
			}
		}
	}
	return "Yes"
}

func main() {
	var s string
	s = readLine()

	ans := check(s)
	fmt.Println(ans)
}
