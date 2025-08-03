package lexer

import (
	"dagon/src/syntax/tokens"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	Int y = 20; 
	Int x = 10;
	Int x1x2 = 200;`
	tests := []struct {
		expectedType    tokens.TokenType
		expectedLiteral string
	}{
		{tokens.INT, "Int"},
		{tokens.IDENT, "y"},
		{tokens.ASSIGN, "="},
		{tokens.INT_LITERAL, "20"},
		{tokens.SEMICOLON, ";"},
		{tokens.INT, "Int"},
		{tokens.IDENT, "x"},
		{tokens.ASSIGN, "="},
		{tokens.INT_LITERAL, "10"},
		{tokens.SEMICOLON, ";"},
		{tokens.INT, "Int"},
		{tokens.IDENT, "x1x2"},
		{tokens.ASSIGN, "="},
		{tokens.INT_LITERAL, "200"},
		{tokens.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Lexeme != tt.expectedLiteral {
			t.Fatalf("test[%d] - lexeme wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Lexeme)
		}
	}
}
