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

	/* Table Driven Test TDT */

	// Anonymous struct
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{
			shape: Rectangle{Width: 12, Height: 6},
			want:  72.0,
		},
		{
			shape: Circle{Radius: 10},
			want:  314.1592653589793,
		},
		{
			shape: Triangle{Length: 12, Height: 6},
			want:  36.0,
		},
	}

	for _, v := range areaTests {
		got := v.shape.Area()
		if got != v.want {
			t.Errorf("%#v got %g want %g", v.shape, got, v.want)
		}
	}

}
