package structmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(123.45, 56.09)
	want := 12.09

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}
