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