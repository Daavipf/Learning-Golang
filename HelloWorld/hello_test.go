package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying Hello to people", func(t *testing.T) {
		got := Hello("Davi", "")
		want := "Hello, Davi"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello World if empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello in spanish", func(t *testing.T) {
		got := Hello("Davi", "Spanish")
		want := "Hola, Davi"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello in french", func(t *testing.T) {
		got := Hello("Davi", "French")
		want := "Bonjour, Davi"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello in Portuguese", func(t *testing.T) {
		got := Hello("Davi", "Portuguese")
		want := "Olá, Davi"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello in unknown language", func(t *testing.T) {
		got := Hello("Davi", "Japanese")
		want := "Hello, Davi"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
