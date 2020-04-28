package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/SouleBA/monkeyintp/lexer"
)

const PROMPT = ">> "

//Start start a REPL process
func Start(in io.Reader, out io.Reader) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Println(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != "EOF"; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
