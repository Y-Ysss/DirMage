package dirmage

import (
	"fmt"
	"regexp"
	"github.com/AlecAivazis/survey/v2"
)

func EditDirectory() {
	fmt.Println("Edit")
	dirsJson := ReadFile(DirectoriesList)
	rePttrn := regexp.MustCompile("{(.|\f|\n|\r|\t)*?}")
	num := 0
	replaceFunc := func(s []byte) []byte {
		num++
		return ([]byte)(fmt.Sprintf("{%d}", num))
	}
	str := rePttrn.ReplaceAllFunc(dirsJson, replaceFunc)
	fmt.Printf("%s", str)
}

func Edit(info *DirInfo) {
	fmt.Println(info)
	name := info.Name
	promptInput := &survey.Input{
		Message: "Name",
		Default: name,
	}
	survey.AskOne(promptInput, &name)

	path := info.Path
	promptInput = &survey.Input{
		Message: "Path",
		Default: path,
	}
	survey.AskOne(promptInput, &path)

	enabled := info.Enabled
	promptConfirm := &survey.Confirm{
	    Message: "Enable ?",
	    Default: enabled,
	}
	survey.AskOne(promptConfirm, &enabled)

	fmt.Println(name, path, enabled)
	// fmt.Println(dirName, dirPath)
	// dirsJson := ReadFile(DirectoriesList)
	// rePttrn := regexp.MustCompile("{(\f|\n|\r|\t)*?\"name\": ??\"" + dirName + "\"(.|\f|\n|\r|\t)*?}")
	// str := rePttrn.Find(dirsJson)
	// fmt.Printf("%s", str)

}