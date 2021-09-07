package arraysslices

import "testing"

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
}
