package main

import "testing"

func TestHello(t *testing.T) {
	/*
		Arrange
		Assert
		Act
	*/

	// arrange
	got := Hello("Chris")
	want := "Hello, Chris"

	// assert
	if got != want {

		// act
		t.Errorf("got %q want %q", got, want)
	}
}
