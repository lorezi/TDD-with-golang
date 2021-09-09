package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"golang": "Go is all about type"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("golang")
		want := "Go is all about type"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("Java")
		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "golang"
		definition := "late have I known golang"

		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "golang"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	word := "golang"
	definition := "Go is all about type"
	newDefinition := "Go forever"

	// add a word to the dictionary
	dictionary := Dictionary{word: definition}
	// update the dictionary
	dictionary.Update(word, newDefinition)
	assertDefinition(t, dictionary, word, newDefinition)

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search("golang")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != definition {
		t.Errorf("got %q, want %q", got, definition)
	}

}
