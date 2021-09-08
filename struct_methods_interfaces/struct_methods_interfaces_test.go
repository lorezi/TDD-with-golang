package structmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.00, 10.00)
	want := 40.00

	if got != want {
		t.Errorf("got %.2f wanted %.2f", got, want)
	}
}
