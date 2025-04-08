package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/TusharAbhinav/monkey/lexer"
	Token "github.com/TusharAbhinav/monkey/token"
)

// REPL stands for Read-Eval-Print Loop
// It is a simple interactive programming environment that takes single user inputs,
// evaluates them, and returns the result to the user.
// The REPL is a common way to interact with programming languages, especially during development and debugging.
var PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != Token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
