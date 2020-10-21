package dirmage

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

func ReadFile(path string) []byte {
	data, readErr := ioutil.ReadFile(filepath.Join(exePath, path))
	if readErr != nil {
		log.Fatal(readErr)
	}
	return data
}

func WriteFile(path string, data []byte) {
	err := ioutil.WriteFile(filepath.Join(exePath, path), data, 0664)
	if err != nil {
		log.Fatal(err)
	}
}
