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
var indexList []int
var conf Config
var working string

func init() {
	checkExistence("config.toml", defaultConfigTomlData)
	checkExistence("directories.json", defaultDirectoriesJsonData)
	readDefaultFiles()
	directoriesFileCopy(conf.Data.DirsFile)
	readDrectoriesJson()
	workingOn()
}

func checkExistence(name string, value string) {
	if _, err := os.Stat(name); err != nil {
		file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if _, err = file.WriteString(value); err != nil {
			log.Fatal(err)
		}
		fmt.Println("File(" + name + ") does not exist -> File is created.")
	}
}

func readDefaultFiles() {
	if _, err := toml.Decode(defaultConfigTomlData, &conf); err != nil {
		log.Fatal(err)
	}
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatal(err)
	}
}

func directoriesFileCopy(fileName string) {
	if _, err := os.Stat(fileName); err == nil {
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
