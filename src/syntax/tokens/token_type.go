package tokens

type TokenType int16

const (
	// Special tokens
	IDENT TokenType = iota
	EOF
	ILLEGAL

	// Operators
	ASSIGN
	PLUS
	MINUS
	TIMES
	DIVIDE

	// Logic operators
	EQ
	NEQ
	LT
	GT
	LE
	GE
	AND
	OR
	NOT

	// Logic values
	TRUE
	FALSE

	// Separators
	SEMICOLON
	COMMA

	// Tipos de variables
	INT
	REAL
	STRING
	BOOLEAN

	//Data
	INT_LITERAL
	REAL_LITERAL
	STRING_LITERAL

	// Structure tokens
	LPAREN   //(
	RPAREN   //)
	LBRACE   //{
	RBRACE   //}
	LBRACKET //[
	RBRACKET //]

	// Keywords
	IF
	ELSE
	FOR
	RETURN
	BREAK
	FUNCTION
)

var tokenNames = map[TokenType]string{
	// Special tokens
	IDENT:   "IDENT",
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",

	// Operators
	ASSIGN: "ASSIGN",
	PLUS:   "PLUS",
	MINUS:  "MINUS",
	TIMES:  "TIMES",
	DIVIDE: "DIVIDE",

	// Logic operators
	EQ:  "EQ",
	NEQ: "NEQ",
	NOT: "NOT",
	LT:  "LT",
	GT:  "GT",
	LE:  "LE",
	GE:  "GE",
	AND: "AND",
	OR:  "OR",

	// Logic values
	TRUE:  "TRUE",
	FALSE: "FALSE",

	// Separators
	SEMICOLON: "SEMICOLON",
	COMMA:     "COMMA",

	// Data types
	INT:     "INT",
	REAL:    "REAL",
	STRING:  "STRING",
	BOOLEAN: "BOOLEAN",

	//Data values
	INT_LITERAL:    "INT_LITERAL",
	REAL_LITERAL:   "REAL_LITERAL",
	STRING_LITERAL: "STRING_LITERAL",

	// Structure tokens
	LPAREN:   "LPAREN",
	RPAREN:   "RPAREN",
	LBRACE:   "LBRACE",
	RBRACE:   "RBRACE",
	LBRACKET: "LBRACKET",
	RBRACKET: "RBRACKET",

	// Keywords
	IF:       "IF",
	ELSE:     "ELSE",
	FOR:      "FOR",
	BREAK:    "BREAK",
	RETURN:   "RETURN",
	FUNCTION: "FUNCTION",
}

var keywords = map[string]TokenType{
	"Int":    INT,
	"Real":   REAL,
	"String": STRING,
	"Bool":   BOOLEAN,

	"if":     IF,
	"else":   ELSE,
	"for":    FOR,
	"return": RETURN,
	"break":  BREAK,
	"func":   FUNCTION,

	"true":  TRUE,
	"false": FALSE,
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
