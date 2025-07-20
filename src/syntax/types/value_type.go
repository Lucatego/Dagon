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
)
