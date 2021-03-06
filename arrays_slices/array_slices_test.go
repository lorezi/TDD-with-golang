package arraysslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		xs := []int{1, 2, 3}
		got := Sum(xs)
		expected := 6

		if got != expected {
			t.Errorf("expected '%d' got '%d' given, %v", expected, got, xs)
		}
	})

	t.Run("collection of varying number of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}
		assertCorrectMessage(t, got, expected)

	})

	t.Run("calculate the totals of the 'tails' of each slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		assertCorrectMessage(t, got, expected)
	})

	t.Run("sum up the tail of an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		expected := []int{0, 9}
		assertCorrectMessage(t, got, expected)
	})
}
