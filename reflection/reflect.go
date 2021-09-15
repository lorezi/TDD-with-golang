package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			fn(field.String())
		}

		if field.Kind() == reflect.Struct {
			// for i := 0; i < field.NumField(); i++ {
			// 	nestedField := field.Field(i)
			// 	if nestedField.Kind() == reflect.String {
			// 		fn(nestedField.String())
			// 	}

			// }
			walk(field.Interface(), fn)
		}

	}
}
