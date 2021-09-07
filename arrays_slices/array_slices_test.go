package arraysslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, expected int, numbers []int) {
		if got != expected {
			t.Errorf("expected '%d' got '%d' given, %v", expected, got, numbers)
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		xs := []int{1, 2, 3}
		got := Sum(xs)
		expected := 6

		assertCorrectMessage(t, got, expected, xs)
	})

	t.Run("collection of varying number of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}

	})
}
