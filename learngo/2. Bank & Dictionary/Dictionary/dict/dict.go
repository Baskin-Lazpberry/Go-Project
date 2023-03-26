package dict

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("not found")
var errWordExists = errors.New("word already exists")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	if err == nil {
		d[word] = def
	} else if err == errNotFound {
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
