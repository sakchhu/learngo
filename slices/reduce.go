package slices

func Reduce(numbers []int) int {
	var sum int

	for _, n := range numbers {
		sum += n
	}

	return sum
}
