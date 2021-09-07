package arraysslices

import "testing"

func TestArrayslice(t *testing.T) {
	xs := [5]int{1, 2, 3, 4, 5}
	got := Sum(xs)
	expected := 15

	if got != expected {
		t.Errorf("expected '%d' got '%d' given, %v", expected, got, xs)
	}

}
