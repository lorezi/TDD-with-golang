package main

import "testing"

func TestHello(t *testing.T) {
	/*
		THREE KEY STRUCTURES FOR UNIT TEST
			Arrange
			Assert
			Act
	*/

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		// act
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		// arrange
		got := Hello("Lorezi", "English")
		want := "Hello, Lorezi"

		// assert
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		// arrange
		got := Hello("", "English")
		want := "Hello, World"
		// assert
		assertCorrectMessage(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Alvaro", "Spanish")
		want := "Hola, Alvaro"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Pogba", "French")
		want := "Bonjour, Pogba"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Pidgin", func(t *testing.T) {
		got := Hello("Akpos", "Pidgin")
		want := "Afa, Akpos"
		assertCorrectMessage(t, got, want)
	})

}
