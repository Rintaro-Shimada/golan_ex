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

func readLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()

	//byte の取得
	white := "."[0]
	black := "#"[0]

	mat := make([]string, h)
	for i := 0; i < h; i++ {
		mat[i] = readLine()
	}

	count := 0

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if mat[i][j] == white {
				continue
			}
			if mat[i-1][j] == white && mat[i][j-1] == white {
				count++
			}
			if mat[i-1][j] == white && mat[i][j+1] == white {
				count++
			}
			if mat[i+1][j] == white && mat[i][j-1] == white {
				count++
			}
			if mat[i+1][j] == white && mat[i][j+1] == white {
				count++
			}
			if mat[i-1][j] == black && mat[i][j-1] == black && mat[i-1][j-1] == white {
				count++
			}
			if mat[i-1][j] == black && mat[i][j+1] == black && mat[i-1][j+1] == white {
				count++
			}
			if mat[i+1][j] == black && mat[i][j-1] == black && mat[i+1][j-1] == white {
				count++
			}
			if mat[i+1][j] == black && mat[i][j+1] == black && mat[i+1][j+1] == white {
				count++
			}
		}
	}
	fmt.Println(count)
}
