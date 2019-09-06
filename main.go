package main

import (
	"fmt"
	"os"
    "os/signal"
)

func main() {
	go func() {
		arglen := len(os.Args)
		if arglen == 1 {
			// CommandExec()
			SelectDirectory()
		} else if arglen == 2{
			arg := os.Args[1]
			switch arg{
			case "edit":
				fmt.Printf("Exec %s\n", arg)
			default:
				fmt.Printf("'%s' is an invalid command line argument.", arg)
			}
		}
	}()

	quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt)
    <-quit
    fmt.Print("\nKeyboard Interrupt (Ctrl+C)\n")
}
