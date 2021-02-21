package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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

func check(n int) int {
	strN := strconv.Itoa(n)
	splitN := strings.Split(strN, "")
	sort.Strings(splitN)

	min := 0
	max := 0
	length := len(splitN)

	for i, e := range splitN {
		num, _ := strconv.Atoi(e)
		min += int(math.Pow(10, float64(length-i-1))) * num
		max += int(math.Pow(10, float64(i))) * num
	}
	return max - min
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	n2 := nextInt()

	var ans int
	for i := 0; i < n2; i++ {
		if i == 0 {
			ans = check(n)
		} else {
			ans = check(ans)
		}
	}

	fmt.Println(ans)
}
