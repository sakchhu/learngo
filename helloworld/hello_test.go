package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("say 'Hello, Club' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, Club"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to the club!", func(t *testing.T) {
		got := Hello("Klein", "")
		want := "Hello, Klein"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello in Spanish", func(t *testing.T) {
		got := Hello("Audrey", spanish)
		want := "Hola, Audrey"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("Alger", french)
		want := "Bonjour, Alger"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello in Nepali", func(t *testing.T) {
		got := Hello("Derrick", nepali)
		want := "Namaste, Derrick"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // inform the test suite that this function is a test helper.
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
