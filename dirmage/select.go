package dirmage

import (
	"fmt"
	"log"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func Select(fn func(*dirInfo), opt ...bool) {
	visible := false
	text := conf.Selector.Text
	if len(opt) == 1 {
		visible = opt[0]
		text = conf.Selector.EditText
	}
	dirs := dirsList[working]
	var dirsNameList []string
	for _, dir := range dirs {
		if dir.Enabled || visible {
			s := dirInfoFormatter(dir, text)
			dirsNameList = append(dirsNameList, s)
			// dirsList[dirInfo] = dir
		}
	}

	var selectDir int
	prompt := &survey.Select{
		Message:  "Choose a directory:",
		Options:  dirsNameList,
		PageSize: 15,
	}

	if surveyErr := survey.AskOne(prompt, &selectDir); surveyErr != nil {
		log.Fatal(surveyErr)
	}

	fmt.Println(dirs[selectDir].Name, dirs[selectDir].Path)

	// if selectDir != -1 {
	fn(&dirs[selectDir])
	// }
}

func dirInfoFormatter(di dirInfo, text string) string {
	enabled := conf.Selector.EnabledText[0]
	if !di.Enabled {
		enabled = conf.Selector.EnabledText[1]
	}
	r := strings.NewReplacer("{$Name}", di.Name, "{$Desc}", di.Description, "{$Path}", di.Path, "{$Enabled}", enabled)
	return r.Replace(text)
}
