package dirmage

import (
	"encoding/json"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func Add() {
	var name, desc, path string
	var enabled bool
	save := true

	survey.AskOne(
		&survey.Input{
			Message: "Name",
		}, &name)

	survey.AskOne(
		&survey.Input{
			Message: "Description",
		}, &desc)

	survey.AskOne(
		&survey.Input{
			Message: "Path",
		}, &path)

	survey.AskOne(
		&survey.Confirm{
			Message: "Enable ?",
		}, &enabled)

	fmt.Println("Name       :", name)
	fmt.Println("Description:", desc)
	fmt.Println("Path       :", path)
	fmt.Println("Enabled    :", enabled)

	survey.AskOne(
		&survey.Confirm{
			Message: "Save ?",
			Default: save,
		}, &save)
	if save {
		dirsList[working] = append(dirsList[working], NewDirInfo(name, desc, path, enabled))
		jsonBytes, err := json.MarshalIndent(dirsList, "", "  ")
		if err != nil {
			fmt.Print(err)
			return
		}
		WriteFile(conf.Data.DirsFile, jsonBytes)
		fmt.Println("Saved")
	} else {
		fmt.Println("Not save")
	}
}
