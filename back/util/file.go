package util

import (
	"io/ioutil"
)

// ReadFile returns the bytes of a file searched in the path and beyond it
func ReadFile(path string) (bytes []byte) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		bytes, err = ioutil.ReadFile("../" + path)
	}

	if err != nil {
		panic(err)
	}

	return bytes
}

// ReadFileErr returns the bytes of a file searched in the path and beyond it
func ReadFileErr(path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		bytes, err = ioutil.ReadFile("../" + path)
	}

	if err != nil {
		return nil, err
	}

	return bytes, nil
}
