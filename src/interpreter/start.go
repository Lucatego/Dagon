package interpreter

import (
	"bufio"
	"dagon/src/core"
	"dagon/src/interpreter/evaluator"
	"dagon/src/syntax/lexer"
	"dagon/src/syntax/parser"
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

		evaluated := evaluator.Eval(program)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
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
