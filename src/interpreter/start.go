package interpreter

import "fmt"

func Start() {
	fmt.Println("Dagon - A basic interpreted language made in Go.")
	for {
		fmt.Print("> ")
		var input string
		fmt.Scanf("%s\n", &input)
		if input == "exit" {
			break
		}
		SyntaxValidator()
		// inputManager()
	}
}
