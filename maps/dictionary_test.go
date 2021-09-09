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
	dictionary := Dictionary{}

	dictionary.Add("golang", "late have I known golang")
	want := "late have I known golang"
	got, err := dictionary.Search("golang")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
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
