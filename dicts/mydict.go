package mydict

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("NOT FOUND")

func (d Dictionary) Search(word string) (string, error) {
	val, exists := d[word]
	if exists {
		return val, nil
	}
	return "", errNotFound
}

func (d Dictionary) AddWord(word string) (string, error) {

}
