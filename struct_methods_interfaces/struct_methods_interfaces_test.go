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
		name    string
		shape   Shape
		hasArea float64
	}{
		{
			name:    "Rectangle",
			shape:   Rectangle{Width: 12, Height: 6},
			hasArea: 72.0,
		},
		{
			name:    "Circle",
			shape:   Circle{Radius: 10},
			hasArea: 314.1592653589793,
		},
		{
			name:    "Triangle",
			shape:   Triangle{Length: 12, Height: 6},
			hasArea: 36.0,
		},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g hasArea %g", tt.shape, got, tt.hasArea)
			}
		})
	}

}
