package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	h, w := nextInt(), nextInt()

	matrix := make(map[int]string, h)
	for i := 0; i < h; i++ {
		x := nextLine()
		matrix[i] = x
	}

	ans := 0

	for i := 0; i < h-1; i++ {
		for j := 0; j < w-1; j++ {
			cnt := 0
			if matrix[i][j] == "#"[0] {
				cnt++
			}
			if matrix[i+1][j] == "#"[0] {
				cnt++
			}
			if matrix[i][j+1] == "#"[0] {
				cnt++
			}
			if matrix[i+1][j+1] == "#"[0] {
				cnt++
			}
			if cnt == 1 || cnt == 3 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
