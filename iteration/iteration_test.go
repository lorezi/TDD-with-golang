package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	iteration := Iteration("a")
	expected := "aaaaa"

	if iteration != expected {
		t.Errorf("expected %q got %q", expected, iteration)
	}
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iteration("a")
	}
}

func ExampleIteration() {
	result := Iteration("a")
	fmt.Println(result)
	// Output: aaaaa
}
