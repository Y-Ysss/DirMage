package dirmage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var dirsList map[string][]dirInfo = make(map[string][]dirInfo)
var conf Config

func Initialize() {
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatal(err)
	}
	fmt.Println(conf)
	if !fileExists(conf.Data.DirsFile) {
		conf.Data.DirsFile = "directories.json"
	}

	readDrectoriesJson()
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func readDrectoriesJson() {
	dirsJson := ReadFile(conf.Data.DirsFile)
	if unmsErr := json.Unmarshal(([]byte)(dirsJson), &dirsList); unmsErr != nil {
		log.Fatal(unmsErr)
	}
}
