package lexer

import (
	"dagon/src/syntax/tokens"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() tokens.Token {
	var tok tokens.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = tokens.Token{Type: tokens.EQ, Lexeme: "=="}
		} else {
			tok = newToken(tokens.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(tokens.SEMICOLON, l.ch)
	case ',':
		tok = newToken(tokens.COMMA, l.ch)
	case '+':
		tok = newToken(tokens.PLUS, l.ch)
	case '-':
		tok = newToken(tokens.MINUS, l.ch)
	case '*':
		tok = newToken(tokens.TIMES, l.ch)
	case '/':
		tok = newToken(tokens.DIVIDE, l.ch)
	case '(':
		tok = newToken(tokens.LPAREN, l.ch)
	case ')':
		tok = newToken(tokens.RPAREN, l.ch)
	case '{':
		tok = newToken(tokens.LBRACE, l.ch)
	case '}':
		tok = newToken(tokens.RBRACE, l.ch)
	case '[':
		tok = newToken(tokens.LBRACKET, l.ch)
	case ']':
		tok = newToken(tokens.RBRACKET, l.ch)
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = tokens.Token{Type: tokens.NEQ, Lexeme: "!="}
		} else {
			tok = newToken(tokens.NOT, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			l.readChar()
			tok = tokens.Token{Type: tokens.LE, Lexeme: "<="}
		} else {
			tok = newToken(tokens.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			l.readChar()
			tok = tokens.Token{Type: tokens.GE, Lexeme: ">="}
		} else {
			tok = newToken(tokens.GT, l.ch)
		}
	case '"':
		tok.Type = tokens.STRING_LITERAL
		tok.Lexeme = l.readString()
		return tok
	case 0:
		tok.Lexeme = ""
		tok.Type = tokens.EOF
	default:
		if isLetter(l.ch) {
			tok.Lexeme = l.readIdentifier()
			tok.Type = tokens.LookupIdent(tok.Lexeme)
			return tok
		} else if isDigit(l.ch) {
			tok.Lexeme, tok.Type = l.readNumber()
			return tok
		} else {
			tok = newToken(tokens.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return tok
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	// Returns an slice, in other words, a substring of the input
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() (string, tokens.TokenType) {
	position := l.position
	tokenType := tokens.INT_LITERAL

	for isDigit(l.ch) {
		l.readChar()
	}

	if l.ch == '.' {
		tokenType = tokens.REAL_LITERAL
		l.readChar()

		for isDigit(l.ch) {
			l.readChar()
		}
	}

	return l.input[position:l.position], tokenType
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readString() string {
	l.readChar()
	position := l.position

	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}

	str := l.input[position:l.position]
	l.readChar()
	return str
}

func newToken(tokenType tokens.TokenType, ch byte) tokens.Token {
	return tokens.Token{Type: tokenType, Lexeme: string(ch)}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
