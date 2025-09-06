package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 0 times returns nil", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		assertCorrectResult(t, repeated, expected)
	})

	repeated := Repeat("x", 5)
	expected := "xxxxx"

	assertCorrectResult(t, repeated, expected)
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 200)
	}
}

func assertCorrectResult(t testing.TB, repeated, expected string) {
	t.Helper()

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("69", 4)
	fmt.Println(repeated)
	// Output: 69696969
}
