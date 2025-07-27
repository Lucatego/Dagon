package interpreter

import (
	"bufio"
	"dagon/src/core"
	"dagon/src/syntax/lexer"
	"dagon/src/syntax/parser"
	"dagon/src/syntax/tokens"
	"fmt"
	"io"
)

func Start(in io.Reader, out io.Writer) {
	presentation()
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf("%s ~ %s > ", core.Username, core.Dir)
		input := scanner.Scan()
		if !input {
			break
		}
		line := scanner.Text()
		if line == "exit" {
			break
		}
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")

		for tok := l.NextToken(); tok.Type != tokens.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
		//tokens := SyntaxValidator(&l)
		//InputManager(tokens)
	}
}

func presentation() {
	fmt.Println("Dagon - A basic interpreted language made in Go - v0.0.1a-pre-alpha")
	fmt.Println("Type 'exit' to quit the interpreter.")
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
