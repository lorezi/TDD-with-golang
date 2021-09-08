package structmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.00, 10.00)
	want := 40.00

	if got != want {
		t.Errorf("got %.2f wanted %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(12.00, 6.0)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f wanted %.2f", got, want)
	}

}
