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

func chmin(ans, p int) int {
	if ans == 0 {
		ans = p
	} else if ans > p {
		ans = p
	}
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()

	var ans int
	for i := 0; i < n; i++ {
		a := nextInt()
		p := nextInt()
		x := nextInt()
		if x > a {
			ans = chmin(ans, p)
		}
	}

	if ans == 0 {
		ans = -1
	}
	fmt.Println(ans)
}
