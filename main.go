package main

import (
	"dagon/src/core"
	"dagon/src/interpreter"
	"dagon/src/memory"
)

func main() {
	// The Initial Allocation of memory for the runtime to work
	memory.InitialAllocation()
	// The Runtime starts - It manages the memory and orchestrates the
	// components
	core.StartRuntime()
	// The CLI, the way the user interacts with the Runtime
	interpreter.Start()
}
