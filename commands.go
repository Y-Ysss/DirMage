package main

import (
	// "os"
	// "os/exec"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"github.com/AlecAivazis/survey/v2"
)

func dirInfoFormatter(di DirInfo) string {
	return fmt.Sprintf("%s (%s)", di.Name, di.Path)
}

func SelectDirectory() {
	dirsJson, readErr := ioutil.ReadFile("directories.json")
	if readErr != nil {
		log.Fatal(readErr)
	}
	rePttrn := regexp.MustCompile("/\\*.*?\\*/|//.*\n")
	dirsJsonStr := rePttrn.ReplaceAllString(string(dirsJson), "")

	var dirs []DirInfo
	if unmsErr := json.Unmarshal(([]byte)(dirsJsonStr), &dirs); unmsErr != nil {
		log.Fatal(unmsErr)
	}

	var dirsList []string
	for _, dir := range dirs {
		if dir.Enabled {
			dirsList = append(dirsList, dirInfoFormatter(dir))
		}
	}

	var selectDir string
	prompt := &survey.Select{
		Message: "Choose a directory:",
		Options: dirsList,
		PageSize: 15,
	}

	if surveyErr := survey.AskOne(prompt, &selectDir); surveyErr != nil {
		log.Fatal(surveyErr)
	}

	if selectDir != "" {
		for _, dir := range dirs {
			if dirInfoFormatter(dir) == selectDir {
				// fmt.Printf("\nName : %s\nPath : %s\n", dir.Name, dir.Path)
				Shell(dir.Name, dir.Path)
			}
		}
	}
}

