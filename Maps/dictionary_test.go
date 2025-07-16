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

		AssertStrings(t, actual, expected)
	})

	t.Run("searching for an unknown word", func(t *testing.T) {
		_, err := dictionary.Search("car")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		AssertStrings(t, err.Error(), ErrNotFound.Error())
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		AssertNoError(t, err)
		AssertDefinition(t, dictionary, word, definition)
	})

	t.Run("add existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		AssertError(t, err, ErrWordExists)
		AssertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definiton for word"

		err := dictionary.Update(word, newDefinition)

		AssertNoError(t, err)
		AssertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("update non-existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		newDefinition := "new definiton for word"

		err := dictionary.Update(word, newDefinition)

		AssertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		AssertError(t, err, ErrNotFound)
	})

	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)

		AssertError(t, err, ErrWordDoesNotExist)
	})
}
