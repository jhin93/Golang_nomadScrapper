package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word] // value는 찾고자 하는 값(d[word])이고 exists는 boolean이다.
	if exists {
		// 찾는 value가 있다면, error 값은 nil로 반환한다.
		//리턴 형태는 (string, error) 이고, 찾는 값이 존재해서 에러가 없으니까 nil로 반환하는 것.
		return value, nil
	}
	return "", errNotFound // 찾는 값이 없으니까 string은 "", 에러는 미리 만들어둔 에러변수를 반환한다.
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}
