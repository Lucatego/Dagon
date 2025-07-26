package tokens

type TokenType int16

const (
	// Single character tokens
	EQUAL = iota
	PLUS
	MINUS
	TIMES
	DIVIDE
	// Separator character tokens
	SEMICOLON
	COMMA
	// Data type tokens
	INTERGER
	REAL
	STRING
)
