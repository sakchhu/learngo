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

func ReduceAllTails(slices ...[]int) []int {
	var sums []int

	for _, slice := range slices {

		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			tail := slice[1:]
			sums = append(sums, Reduce(tail))
		}
	}

	return sums
}
