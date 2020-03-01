package dirmage

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"github.com/AlecAivazis/survey/v2"
)

func SelectDirectory(fn func(*DirInfo)) {
	// dirsJson, readErr := ioutil.ReadFile(DirectoriesList)
	// if readErr != nil {
	// 	log.Fatal(readErr)
	// }
	dirsJson := ReadFile(DirectoriesList)
	rePttrn := regexp.MustCompile("/\\*.*?\\*/|//.*\n")
	dirsJsonStr := rePttrn.ReplaceAllString(string(dirsJson), "")

	var dirs []DirInfo
	if unmsErr := json.Unmarshal(([]byte)(dirsJsonStr), &dirs); unmsErr != nil {
		log.Fatal(unmsErr)
	}
	var dirsNameList []string
	DirsList = make(map[string]DirInfo)
	for _, dir := range dirs {
		if dir.Enabled {
			dirInfo := dirInfoFormatter(dir)
			dirsNameList = append(dirsNameList, dirInfo)
			DirsList[dirInfo] = dir
		}
	}

	var selectDir string
	prompt := &survey.Select{
		Message: "Choose a directory:",
		Options: dirsNameList,
		PageSize: 15,
	}

	if surveyErr := survey.AskOne(prompt, &selectDir); surveyErr != nil {
		log.Fatal(surveyErr)
	}

	if selectDir != "" {
		info := DirsList[selectDir]
		fn(&info)
	}
}

func dirInfoFormatter(di DirInfo) string {
	return fmt.Sprintf("%s (%s)", di.Name, di.Path)
}