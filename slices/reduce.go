package slices

// Reduce returns the sum all the integers in the given `[]int` slice.
func Reduce(numbers []int) int {
	var sum int

	for _, n := range numbers {
		sum += n
	}

	return sum
}

func ReduceAll(slices ...[]int) []int {
	var sums []int

	for _, slice := range slices {
		sums = append(sums, Reduce(slice))
	}

	return sums
}
