package main

import (
	"fmt"
	"os"

	"github.com/Rayan-BA/monkey/repl"
)

func main() {
	fmt.Print("Hello! This is the Monkey programming language!\n")
	repl.Start(os.Stdin, os.Stdout)
}
