package main

import (
	"fmt"
	"github.com/TusharAbhinav/monkey/repl"
	"os"
)

func main() {
	fmt.Println("Hello, Monkey!")
	fmt.Println("This is the Monkey programming language!")
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
