package types

/*
ValueType:

It works as an enum to allocate the variables.
*/
type ValueType uint8

const (
	IntergerType ValueType = iota
	RealType
	StringType
	// COUNTER_ValueType is used to count the number of value types
	COUNTER_ValueType
)
