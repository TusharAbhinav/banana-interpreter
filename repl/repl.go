package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/TusharAbhinav/monkey/lexer"
	"github.com/TusharAbhinav/monkey/parser"
)

// REPL stands for Read-Eval-Print Loop
// It is a simple interactive programming environment that takes single user inputs,
// evaluates them, and returns the result to the user.
// The REPL is a common way to interact with programming languages, especially during development and debugging.
var PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

const MONKEY_FACE = `
.---.  .-"     "-.  .---.
      / .-. \/  .-. .-.  \/ .-. \
     | / '-.))  / Y \  ((.-' \ |
     |/ /   \|  \0 0/  |/   \ \|
      |  .-. \.-"""""""-./ .-.  |
      | /   '-/    ^    \-'   \ |
      |/-''-\|   /_\   |/-''-\|
      /  / \  \ /___\ /  / \  \
      '._ '._'.'-._.-'.'_.' _.'
         '-._ ''====='' _.-'
             ''--._.--''
                __,__
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
