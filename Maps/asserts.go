package maps

import "testing"

func AssertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func AssertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func AssertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	actual, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("unexpected error: ", err)
	}

	AssertStrings(t, actual, definition)
}

func AssertStrings(t testing.TB, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("actual %q expected %q", actual, expected)
	}
}
