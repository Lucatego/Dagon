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
	ASSIGN
	PLUS
	MINUS
	TIMES
	DIVIDE

	EQ
	NEQ
	NEG
	LT
	GT
	LE
	GE
	AND
	OR
	NOT

	// Separadores
	SEMICOLON
	COMMA

	// Tipos de variables
	INT_TYPE
	REAL_TYPE
	STRING_TYPE

	// Mixtos
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACKET
	RBRACKET

	LET
	IF
	ELSE
	WHILE
	FOR
	RETURN
	FUNCTION
	TRUE
	FALSE
)

var tokenNames = map[TokenType]string{
	INTEGER_LITERAL: "INTEGER_LITERAL",
	REAL_LITERAL:    "REAL_LITERAL",
	STRING_LITERAL:  "STRING_LITERAL",
	IDENT:           "IDENT",
	EOF:             "EOF",
	ILLEGAL:         "ILLEGAL",
	ASSIGN:          "ASSIGN",
	PLUS:            "PLUS",
	MINUS:           "MINUS",
	TIMES:           "TIMES",
	DIVIDE:          "DIVIDE",
	EQ:              "EQ",
	NEQ:             "NEQ",
	NEG:             "NEG",
	LT:              "LT",
	GT:              "GT",
	LE:              "LE",
	GE:              "GE",
	AND:             "AND",
	OR:              "OR",
	NOT:             "NOT",
	SEMICOLON:       "SEMICOLON",
	COMMA:           "COMMA",
	INT_TYPE:        "INT_TYPE",
	REAL_TYPE:       "REAL_TYPE",
	STRING_TYPE:     "STRING_TYPE",
	LPAREN:          "(",
	RPAREN:          ")",
	LBRACE:          "{",
	RBRACE:          "}",
	LBRACKET:        "[",
	RBRACKET:        "]",
	LET:             "LET",
	IF:              "IF",
	ELSE:            "ELSE",
	WHILE:           "WHILE",
	FOR:             "FOR",
	RETURN:          "RETURN",
	FUNCTION:        "FUNCTION",
	TRUE:            "TRUE",
	FALSE:           "FALSE",
}

var keywords = map[string]TokenType{
	"Int":    INT_TYPE,
	"Real":   REAL_TYPE,
	"String": STRING_TYPE,
	"if":     IF,
	"else":   ELSE,
	"for":    FOR,
	"while":  WHILE,
	"return": RETURN,
	"fnt":    FUNCTION,
	"true":   TRUE,
	"false":  FALSE,
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
