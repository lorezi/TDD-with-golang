package main

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

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		r := Rectangle{12, 6}
		want := 72.0
		checkArea(t, r, want)
	})

	t.Run("circles", func(t *testing.T) {
		c := Circle{10}
		want := 314.1592653589793

		checkArea(t, c, want)

	})

}
