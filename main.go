package main

import (
	"fmt"
	"os"
    "os/signal"
    "DirMage/dirmage"
)

func main() {
	go func()  {
		dirmage.Initialize()
		arglen := len(os.Args)
		if arglen == 1 {
			dirmage.SelectDirectory(dirmage.Shell)
		} else if arglen == 2{
			arg := os.Args[1]
			switch arg{
			case "add":
				fmt.Printf("Add %s\n", arg)
			case "edit":
				dirmage.SelectDirectory(dirmage.Edit)
				// dirmage.EditDirectory()
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
