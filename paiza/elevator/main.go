package main

import "fmt"

func main() {
	var n, n2 int
	fmt.Scan(&n)
	floor := 1
	count :=0
	for i:=0;i< n; i++{
		if i == 0{
			fmt.Scan(&n2)
			count += n2 - floor
			floor = n2
		} else {
			fmt.Scan(&n2)
			if n2 > floor{
				count += n2 - floor
				floor = n2
			} else {
				count += floor - n2
				floor = n2
			}
		}
	}
	fmt.Println(count)
}
