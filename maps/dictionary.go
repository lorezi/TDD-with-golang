package main

type Dictionary map[string]string

func (d Dictionary) Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
