package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile
}

type Profile struct {
	Age  int
	City string
}

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

		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Lawrence", "Lagos"},
			[]string{"Lawrence", "Lagos"},
		},

		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Lawrence", 23},
			[]string{"Lawrence"},
		},

		{
			"Nested fields",
			Person{
				"Lawrence",
				Profile{
					23, "Lagos",
				},
			},
			[]string{"Lawrence", "Lagos"},
		},

		{
			"Pointers to things",
			&Person{
				"Lawrence",
				Profile{
					23, "Lagos",
				},
			},
			[]string{"Lawrence", "Lagos"},
		},

		{
			"Slices",
			[]Profile{
				{33, "Lagos"},
				{23, "Abuja"},
			},
			[]string{"Lagos", "Abuja"},
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
