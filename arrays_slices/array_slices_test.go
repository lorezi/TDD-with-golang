package arraysslices

import "testing"

func TestArray(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	expected := 15

	if got != expected {
		t.Errorf("expected '%d' got '%d' given, %v", expected, got, numbers)
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
