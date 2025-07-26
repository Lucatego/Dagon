package memory

import (
	"dagon/src/syntax/types"
)

var (
	accessVariables [](map[string]any)
)

func InitialAllocation() {
	// Create a map for each variable type
	// IntergerType, RealType, StringType
	accessVariables = make([]map[string]any, types.COUNTER_ValueType)
	// Initialize each map
	initializeVariablesMemory()
}

// It creates a map for each variable type. The maps are used to store the data
// of the varaibles in memory. The index is the variable name.
func initializeVariablesMemory() {
	accessVariables[types.IntergerType] = make(map[string]any) // IntergerType
	accessVariables[types.RealType] = make(map[string]any)     // RealType
	accessVariables[types.StringType] = make(map[string]any)   // StringType
}
