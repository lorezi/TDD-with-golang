package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	testCases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Lawrence"},
			[]string{"Lawrence"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.Name, func(t *testing.T) {
			var got []string
			walk(tC.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, tC.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, tC.ExpectedCalls)
			}

		})
	}
}
