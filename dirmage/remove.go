package dirmage

import (
	"encoding/json"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func Remove() {
	Select(OpenRemove, true)
}

func OpenRemove(info *dirInfo, index int) {
	fmt.Println("Name        :", info.Name)
	fmt.Println("Description :", info.Description)
	fmt.Println("Path        :", info.Path)
	fmt.Println("Enabled     :", info.Enabled)

	save := false
	survey.AskOne(
		&survey.Confirm{
			Message: "Remove ?",
			Default: save,
		}, &save)

	if save {
		dirsList[working] = append(dirsList[working][:index], dirsList[working][index+1:]...)
		jsonBytes, err := json.MarshalIndent(dirsList, "", "  ")
		if err != nil {
			fmt.Print(err)
			return
		}
		WriteFile(conf.Data.DirsFile, jsonBytes)
		fmt.Println("Removed")
	} else {
		fmt.Println("No change")
	}
}
