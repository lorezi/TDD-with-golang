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
		definition := "late have I known golang"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "golang has practically helped me to understand those theoretical concepts in cs")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
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
		t.Errorf("got error %q want %q", got, want)
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
