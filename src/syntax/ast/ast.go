package ast

import (
	"dagon/src/syntax/tokens"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type ReturnStatement struct {
	Token       tokens.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Lexeme }

type AssignStatement struct {
	Token tokens.Token
	Name  *Identifier
	Value Expression
}

func (ls *AssignStatement) statementNode()       {}
func (ls *AssignStatement) TokenLiteral() string { return ls.Token.Lexeme }

type Identifier struct {
	Token tokens.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Lexeme }

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
