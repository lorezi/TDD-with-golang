package main

import (
	"testing"
)

func TestSearch(t *testing.T) {

	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"golang": "Go is about type"}

		got := dictionary.Search("golang")
		want := "Go is about type"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"golang": "Go is about type"}

		got := dictionary.Search("java")
		want := "does not exist 🥵"

		assertStrings(t, got, want)
	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
