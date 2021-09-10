package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Lawrence")

	got := buffer.String()
	want := "Hello, Lawrence"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
