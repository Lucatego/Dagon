package tokens

type Token struct {
	// What type of token it is
	Type TokenType
	// The actual string that was parsed
	Lexeme string
}
