package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/prdpx7/cheena/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is cheena programming language!\n", user.Username)
	fmt.Println("Starting REPL...")
	repl.Start(os.Stdin, os.Stdout)
}
