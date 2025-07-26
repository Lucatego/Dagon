package tokens

type Token struct {
	// What type of token it is
	Type TokenType
	// The value of the token, if applicable
	Value any
	// The actual string that was parsed
	Lexeme string
}
