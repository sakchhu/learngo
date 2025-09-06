package slices

// Reduce returns the sum all the integers in the given `[]int` slice.
func Reduce(numbers []int) int {
	var sum int

	for _, n := range numbers {
		sum += n
	}

	return sum
}
