package dirmage

import (
	"os"
	"log"
	"github.com/BurntSushi/toml"
)

func Initialize() {
	var conf Config
    if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
    	log.Fatal(err)
	}
    if(exists(conf.Selector.DirsListPath)) {
    	DirectoriesList = conf.Selector.DirsListPath
    }
    PromptString = conf.Prompt.String
}

func exists(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}