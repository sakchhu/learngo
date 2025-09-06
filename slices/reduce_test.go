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

func TestReduceAllTails(t *testing.T) {

	checkSums := func(t testing.TB, want, got []int) {
		t.Helper()
		if !slices.Equal(want, got) {
			t.Errorf("want %v got %v", want, got)
		}
	}
	t.Run("reduce the tails of multiple slices", func(t *testing.T) {
		got := ReduceAllTails([]int{-1, 1, 20}, []int{30, -40, 10})
		want := []int{21, -30}
		checkSums(t, want, got)
	})
	t.Run("reduce the tails of empty slices safely", func(t *testing.T) {
		got := ReduceAllTails([]int{}, []int{})
		want := []int{0, 0}
		checkSums(t, want, got)
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

func ExampleReduceAllTails() {
	reducedList := ReduceAllTails([]int{20, -1}, []int{43, 200, 30})
	fmt.Println(reducedList)
	// Output: [-1 230]
}
