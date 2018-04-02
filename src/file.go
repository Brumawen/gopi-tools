package gopitools

import (
	"io/ioutil"
)

func ReadAllText(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
