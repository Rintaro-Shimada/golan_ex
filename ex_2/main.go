package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	vexA, vexB := Reader(num)
	Orthogonality(vexA, vexB, num)
}

func Reader(num int) (vexA, vexB []int) {
	for i := 0; i < num; i++ {
		num := 0
		fmt.Scan(&num)
		vexA = append(vexA, num)
	}

	for i := 0; i < num; i++ {
		num := 0
		fmt.Scan(&num)
		vexB = append(vexB, num)
	}
	return vexA, vexB
}

func Orthogonality(vexA, vexB []int, num int) {
	var ans int
	for i := 0; i < num; i++ {
		ans += vexA[i] * vexB[i]
		fmt.Println(ans)
	}
	if ans == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
