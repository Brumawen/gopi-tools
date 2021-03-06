package gopitools

import (
	"io/ioutil"
)

// ReadAllText reads the text from the specified file path
func ReadAllText(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// WriteAllText writes the text to the specified file path
//func WriteAllText(path string, text string) error {
//	return ioutil.WriteFile(path, text, "0666")
//}
