package main

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, exist := d[word]
	if !exist {
		err := errors.New("could not find the word you were looking for")
		return "", err
	}
	return definition, nil
}
