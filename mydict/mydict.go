package mydict

import (
	"errors"
)

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	//if err == errNotFound {
	//	d[word] = def
	//} else if err == nil {
	//	return errWordExists
	//}
	//return nil

	switch {
	case errors.Is(err, errNotFound):
		d[word] = def
	case err == nil:
		return errWordExists
	}
	return nil
}
