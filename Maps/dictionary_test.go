package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("searching for a known word", func(t *testing.T) {
		actual, err := dictionary.Search("test")
		expected := "this is just a test"

		if err != nil {
			t.Fatal("unexpected error")
		}

		assertStrings(t, actual, expected)
	})

	t.Run("searching for an unknown word", func(t *testing.T) {
		_, err := dictionary.Search("car")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err, ErrNotFound)
	})
}

func assertStrings(t testing.TB, actual, expected error) {
	t.Helper()

	if actual != expected {
		t.Errorf("actual %q expected %q", actual, expected)
	}
}
