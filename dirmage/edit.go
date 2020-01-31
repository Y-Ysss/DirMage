package dirmage

import (
	"fmt"
	"regexp"
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

func Edit(dirName, dirPath string) {
	fmt.Println(dirName, dirPath)
	dirsJson := ReadFile(DirectoriesList)
	rePttrn := regexp.MustCompile("{(.|\f|\n|\r|\t)*?\"name\": ??\"" + dirName + "\"(.|\f|\n|\r|\t)*?}")
	str := rePttrn.Find(dirsJson)
	fmt.Printf("%s", str)
}