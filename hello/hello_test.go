package main

import "testing"

func TestHello(t *testing.T) {
	/*
		Arrange
		Assert
		Act
	*/

	// arrange
	got := Hello()
	want := "Hello, world!"

	// assert
	if got != want {

		// act
		t.Errorf("got %q want %q", got, want)
	}
}
