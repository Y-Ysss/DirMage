package dirmage

import (
	"io/ioutil"
	"log"
)

func ReadFile(path string) []byte {
	data, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return data
}

func WriteFile(path string, data []byte) {
	err := ioutil.WriteFile(path, data, 0664)
	if err != nil {
		log.Fatal(err)
	}
}
