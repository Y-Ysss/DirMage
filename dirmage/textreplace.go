package dirmage

import (
	"fmt"
	"regexp"
)

func replaceColor(str string) string {
	rePttrn := regexp.MustCompile("{([0-9]|10|(2|3|4|9|10)[0-7](;(3|4|9|10)[0-7])*)}")
	replaceFunc := func(s string) string {
		return fmt.Sprintf("\x1b[%sm", rePttrn.FindStringSubmatch(s)[1])
	}
	return rePttrn.ReplaceAllStringFunc(str, replaceFunc)
}
