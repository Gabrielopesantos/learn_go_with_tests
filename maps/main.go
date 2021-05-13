package main

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("word not in dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(value string) (string, error) {
	if val, ok := d[value]; ok {
		return val, nil
	}

	return "", ErrNotFound
}

func (d *Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		(*d)[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// func (d *Dictionary) Add(key, value string) error {
// 	if _, ok := (*d)[key]; ok {
// 		return ErrWordExists
// 	}
// 	(*d)[key] = value
// 	return nil
// }

func (d *Dictionary) Update(word, definition string) error {
	if _, ok := (*d)[word]; !ok {
		return ErrWordDoesNotExist
	}

	(*d)[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func Search(dict map[string]string, value string) string {
	return dict[value]
}

func main() {
}
