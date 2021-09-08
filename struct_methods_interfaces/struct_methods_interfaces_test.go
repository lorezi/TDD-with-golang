package structmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{
		Width:  10.00,
		Height: 10.00,
	}
	got := Perimeter(r.Width, r.Height)
	want := 40.00

	if got != want {
		t.Errorf("got %g wanted %g", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		r := Rectangle{12, 6}
		got := Area(r)
		want := 72.0

		if got != want {
			t.Errorf("got %g wanted %g", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		c := Circle{10}
		got := Area(c)
		want := 314.15926

		if got != want {
			t.Errorf("got %g wanted %g", got, want)
		}
	})

}
