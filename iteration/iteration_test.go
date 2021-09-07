package iteration

import "testing"

func TestIteration(t *testing.T) {
	iteration := Iteration("a")
	expected := "aaaaa"

	if iteration != expected {
		t.Errorf("expected %q got %q", expected, iteration)
	}
}
