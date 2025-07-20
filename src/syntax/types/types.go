package types

/*
Types interface:
The types that will be accepted by the interpreter.

- Int = int64

- Real = float64

- String = string
*/
type Types interface {
	int64 | float64 | string
}
