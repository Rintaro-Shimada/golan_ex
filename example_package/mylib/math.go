package mylib

func average(i []int) int {
	total := 0
	for _, i := range i {
		total += i
	}

	return int(total / len(i))
}
