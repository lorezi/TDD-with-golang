package arraysslices

import "testing"

func TestArray(t *testing.T) {
	xs := [5]int{1, 2, 3, 4, 5}
	got := Sum(xs)
	expected := 15

	if got != expected {
		t.Errorf("expected '%d' got '%d' given, %v", expected, got, xs)
	}

}

func TestSlice(t *testing.T) {

	xs := []int{1, 2, 3}
	got := SumN(xs)
	expected := 6
	if got != expected {
		t.Errorf("expected '%d' got '%d' given, %v", expected, got, xs)
	}

}
