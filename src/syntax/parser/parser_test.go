package parser

import (
	"dagon/src/syntax/ast"
	"dagon/src/syntax/lexer"
	"testing"
)

func TestAsignStatements(t *testing.T) {
	input := `
	Int x = 5;
	Int y = 10;
	Int  481923;`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("Program() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testAsignStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testAsignStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "Int" {
		t.Errorf("s.TokenLiteral() not 'Int'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.AssignStatement)
	if !ok {
		t.Errorf("s not *ast.Statement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %s. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not %s. got=%s", name, letStmt.Name)
		return false
	}

	return true
}
