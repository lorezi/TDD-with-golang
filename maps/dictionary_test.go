package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"golang": "Go is about type"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("golang")
		want := "Go is about type"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("Java")
		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertStrings(t, err.Error(), ErrNotFound.Error())
	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
