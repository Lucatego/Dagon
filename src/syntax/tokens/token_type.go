package tokens

type TokenType int16

const (
	// Literales
	INTEGER_LITERAL TokenType = iota
	REAL_LITERAL
	STRING_LITERAL

	// Identificadores y fin
	IDENT
	EOF
	ILLEGAL

	// Operadores
	EQUAL
	PLUS
	MINUS
	TIMES
	DIVIDE

	// Separadores
	SEMICOLON
	COMMA

	// Tipos de variables
	INT_TYPE
	REAL_TYPE
	STRING_TYPE
)

var tokenNames = map[TokenType]string{
	INTEGER_LITERAL: "INTEGER_LITERAL",
	REAL_LITERAL:    "REAL_LITERAL",
	STRING_LITERAL:  "STRING_LITERAL",
	IDENT:           "IDENT",
	EOF:             "EOF",
	ILLEGAL:         "ILLEGAL",
	EQUAL:           "EQUAL",
	PLUS:            "PLUS",
	MINUS:           "MINUS",
	TIMES:           "TIMES",
	DIVIDE:          "DIVIDE",
	SEMICOLON:       "SEMICOLON",
	COMMA:           "COMMA",
	INT_TYPE:        "INT_TYPE",
	REAL_TYPE:       "REAL_TYPE",
	STRING_TYPE:     "STRING_TYPE",
}

var keywords = map[string]TokenType{
	"Int":    INT_TYPE,
	"Real":   REAL_TYPE,
	"String": STRING_TYPE,
}

func (tt TokenType) String() string {
	if name, ok := tokenNames[tt]; ok {
		return name
	}
	return "UNKNOWN"
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
