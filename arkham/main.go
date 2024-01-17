package main

import (
	"fmt"
	"os"
	"os/user"
	"pa-oshea/arkham/repl"
)

// TODO: token add line number and column number for better errors
// TODO: unicode
func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Arkham programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
