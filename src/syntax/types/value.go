package types

/*
	Value struct:
Is the variable that will be defined by the user.

	For example:

When user does:

 > Real pi = 3.141592;

The interpreter will save Value{Type: float64, Data: 3.141592}. This will be
allocated in the runtime memory.

*/
type Value struct {
	Type ValueType
	Data any
}
