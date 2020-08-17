package dirmage

import (
	"encoding/json"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func Edit() {
	Select(OpenEdit, true)
}

func OpenEdit(info *dirInfo, _ int) {
	name := info.Name
	desc := info.Description
	path := info.Path
	enabled := info.Enabled
	save := true

	survey.AskOne(
		&survey.Input{
			Message: "Name",
			Default: name,
		}, &name)

	survey.AskOne(
		&survey.Input{
			Message: "Description",
			Default: desc,
		}, &desc)

	survey.AskOne(
		&survey.Input{
			Message: "Path",
			Default: path,
		}, &path)

	survey.AskOne(
		&survey.Confirm{
			Message: "Enable ?",
			Default: enabled,
		}, &enabled)

	survey.AskOne(
		&survey.Confirm{
			Message: "Save ?",
			Default: save,
		}, &save)

	if save {
		fmt.Println(name, desc, path, enabled)
		info.SetValues(name, desc, path, enabled)
		jsonBytes, err := json.MarshalIndent(dirsList, "", "  ")
		if err != nil {
			fmt.Print(err)
			return
		}
		WriteFile(conf.Data.DirsFile, jsonBytes)
		fmt.Println("Saved")
	} else {
		fmt.Println("No change")
	}
}
