package main

import (
	"fmt"
	"os"
	"os/user"

	"monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! this is the Monkey programming language!\n", user.Username)

	fmt.Println("feel free to input commands")

	repl.Start(os.Stdin, os.Stdout)
}
