package structmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{
		Width:  10.00,
		Height: 10.00,
	}
	got := r.Perimeter(10.00, 10.00)
	want := 40.00

	if got != want {
		t.Errorf("got %.2f wanted %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	r := Rectangle{
		Width:  12.00,
		Height: 6.0,
	}
	got := r.Area(12.00, 6.0)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f wanted %.2f", got, want)
	}

}
