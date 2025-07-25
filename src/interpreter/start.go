package interpreter

import (
	"dagon/src/core"
	"fmt"
)

func Start() {
	fmt.Println("Dagon - A basic interpreted language made in Go.")
	fmt.Printf("Hi %s. Working directory: %s\n", core.Username, core.Dir)
	fmt.Println("Type 'exit' to quit the interpreter.")
	for {
		fmt.Printf("> ")
		var input string
		fmt.Scanf("%s\n", &input)
		if input == "exit" {
			break
		}
		SyntaxValidator()
		// inputManager()
	}
}
