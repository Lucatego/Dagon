package interpreter

import (
	"dagon/src/core"
	"fmt"
)

func Start() {
	presentation()
	for {
		fmt.Printf("%s ~ %s > ", core.Username, core.Dir)
		var input string
		fmt.Scanf("%s\n", &input)
		if input == "exit" {
			break
		}
		tokens := SyntaxValidator(&input)
		InputManager(tokens)
	}
}

func presentation() {
	fmt.Println("Dagon - A basic interpreted language made in Go - v0.0.1a-pre-alpha")
	fmt.Println("Type 'exit' to quit the interpreter.")
}
