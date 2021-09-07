package main

import "testing"

func TestHello(t *testing.T) {
	/*
		THREE KEY PATTERNS FOR UNIT TEST
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
		got := Hello("Lorezi", "")
		want := "Hello, Lorezi"

		// assert
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {

		// arrange
		got := Hello("", "")
		want := "Hello, World"

		// assert
		assertCorrectMessage(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

}
