package memory

import "dagon/src/syntax/types"

var (
	variables map[string]types.Value
)

func InitialAllocation() {
	// Initialize the variables memory
	variables = make(map[string]types.Value)
}
