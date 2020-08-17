package dirmage

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
)

var dirsList map[string][]dirInfo = make(map[string][]dirInfo)
var conf Config
var working string

func init() {
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatal(err)
	}
	directoriesFileCopy(conf.Data.DirsFile)
	readDrectoriesJson()
	workingOn()
}

func workingOn() {
	osTarget := runtime.GOOS
	var osList map[string][]string = make(map[string][]string)
	osList["windows"] = []string{"Windows", "windows", "Win", "win"}
	osList["darwin"] = []string{"MacOS", "darwin", "MacOs", "macos", "Mac", "mac"}
	osList["linux"] = []string{"Linux", "linux"}
	for _, name := range osList[osTarget] {
		_, exsist := dirsList[name]
		if exsist {
			working = name
			return
		}
	}
	fmt.Println("\""+osList[osTarget][0]+"\"", "key not found in", conf.Data.DirsFile)
	fmt.Println("Use the key shown below", osList[osTarget])
	os.Exit(1)
}

func directoriesFileCopy(fileName string) {
	_, err := os.Stat(fileName)
	if err == nil {
		return
	}
	src, err := os.Open("directories.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer src.Close()
	dst, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer dst.Close()
	io.Copy(dst, src)
	dst.Close()
}

func readDrectoriesJson() {
	dirsJson := ReadFile(conf.Data.DirsFile)
	if unmsErr := json.Unmarshal(dirsJson, &dirsList); unmsErr != nil {
		log.Fatal(unmsErr)
		os.Exit(1)
	}
}
