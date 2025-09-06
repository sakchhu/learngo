package slices

import (
	"fmt"
	"slices"
	"testing"
)

func TestReduce(t *testing.T) {

	t.Run("reduce a slice of arbitrary size", func(t *testing.T) {

		numbers := []int{1, 10, 20, 30, -50}

		got := Reduce(numbers)
		want := 11

		if got != want {
			t.Errorf("want %d got %d supplied %v", want, got, numbers)
		}
	})

}

func TestReduceAll(t *testing.T) {

	t.Run("reduce multiple slices, and return a slice containing results for each", func(t *testing.T) {
		positives := []int{30, 40, 50, 60}
		negatives := []int{-10, -20, -30}

		got := ReduceAll(positives, negatives)
		want := []int{180, -60}

		if !slices.Equal(got, want) {
			t.Errorf("want %v got %v", want, got)
		}
	})
}

func ExampleReduce() {
	numbers := []int{1, 121, 10, -100}
	reduced := Reduce(numbers)
	fmt.Println(reduced)
	// Output: 32
}

func ExampleReduceAll() {
	reducedList := ReduceAll([]int{30, -10, 20}, []int{400, 89, 11})
	fmt.Println(reducedList)
	// Output: [40 500]
}
