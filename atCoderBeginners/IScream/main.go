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

func main() {
	sc.Split(bufio.ScanWords)

	A := nextInt()
	B := nextInt()

	fmt.Println(check(A,B))
}

func check(A,B  int) int {
	//アイスクリーム
	if A+B >= 15 && B >= 8{
		return 1
	}

	//アイスミルク
	if A+B >= 10 && B >= 3{
		return 2
	}

	//ラクトアイス
	if A+B >= 3{
		return 3
	}
	return 4
}
