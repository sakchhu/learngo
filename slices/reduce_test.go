package slices

import (
	"fmt"
	"testing"
)

func TestReduce(t *testing.T) {

	t.Run("reduce an array of arbitrary size", func(t *testing.T) {

		numbers := []int{1, 10, 20, 30, -50}

		got := Reduce(numbers)
		want := 11

		if got != want {
			t.Errorf("want %d got %d supplied %v", want, got, numbers)
		}
	})

}

func ExampleReduce() {
	numbers := []int{1, 121, 10, -100}
	reduced := Reduce(numbers)
	fmt.Println(reduced)
	// Output: 32
}
