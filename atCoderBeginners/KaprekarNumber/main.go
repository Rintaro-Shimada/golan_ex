package main

import (
	"bufio"
	"fmt"
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

func ascending(i int) int {
	strNum := strconv.Itoa(i)
	splitNum := strings.Split(strNum, "")
	sort.Sort(sort.Reverse(sort.StringSlice(splitNum)))
	ans, _ := strconv.Atoi(strings.Join(splitNum, ""))
	return ans
}

func descending(i int) int {
	strNum := strconv.Itoa(i)
	splitNum := strings.Split(strNum, "")
	sort.Strings(splitNum)
	ans, _ := strconv.Atoi(strings.Join(splitNum, ""))
	return ans
}

func check(n int) int {
	return ascending(n) - descending(n)
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	n2 := nextInt()

	for i := 0; i < n2; i++ {
		n = check(n)
	}
	fmt.Println(n)
}
