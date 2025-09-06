package slices

import "testing"

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
