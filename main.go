package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/prdpx7/monkey-iptr/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is monkey-iptr programming language!\n", user.Username)
	fmt.Println("Starting REPL...")
	repl.Start(os.Stdin, os.Stdout)
}
