package ast

import (
	"dagon/src/syntax/tokens"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&AssignStatement{
				Token: tokens.Token{Type: tokens.INT_TYPE, Lexeme: "Int"},
				Name: &Identifier{
					Token: tokens.Token{Type: tokens.IDENT, Lexeme: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: tokens.Token{Type: tokens.IDENT, Lexeme: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	expected := "Int myVar = anotherVar;"
	if program.String() != expected {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
