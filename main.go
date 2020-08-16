package main

import (
	"DirMage/dirmage"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	go func() {
		dirmage.Initialize()
		arglen := len(os.Args)
		if arglen == 1 {
			dirmage.Shell()
		} else if arglen == 2 {
			arg := os.Args[1]
			switch arg {
			case "add":
				fmt.Printf("Add %s\n", arg)
			case "edit":
				dirmage.Edit()
			default:
				fmt.Printf("'%s' is an invalid command line argument.", arg)
			}
		}
		os.Exit(0)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Print("\nKeyboard Interrupt (Ctrl+C)\n")
}
