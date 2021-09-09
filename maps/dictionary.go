package main

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add new word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, exist := d[word]
	if !exist {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	// I think, using the switch statement creates redundancy 🤔🤔🤷🏽‍♂️
	// switch err {
	// case ErrNotFound:
	// 	d[word] = definition
	// case nil:
	// 	return ErrWordExists
	// default:
	// 	return nil
	// }

	// return nil

	if err == ErrNotFound {
		d[word] = definition
		return nil
	}

	return ErrWordExists

}

func (d Dictionary) Update(word, definition string) {

}
