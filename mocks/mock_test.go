package main

import (
	"bytes"
	"testing"
)

// spy sleeper
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	iteration := 4

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != iteration {
		t.Errorf("not enough calls to sleeper, want %d got %d", iteration, spySleeper.Calls)
	}
}
