package main

import (
	"os"
	"fmt"
	"os/user"
	"github.com/joalberto902/monkey_interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil { 
		panic(err)
	}

	fmt.Printf("Hello %s! This is the monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
