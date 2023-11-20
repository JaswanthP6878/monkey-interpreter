package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestParser(t *testing.T) {

	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program != nil {
		t.Fatalf("ParserProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Program failed to produce 3 statements:  produced only %d ", len(program.Statements))

	}

	tests := []struct {
		expectedIndentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatment(t, stmt, tt.expectedIndentifier) {
			return
		}
	}

}

func testLetStatment(t *testing.T, s ast.Statement, name string) bool {

	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letstmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement . got %T", s)
		return false
	}

	if letstmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %s . got %s ", name, letstmt.Name.Value)
		return false
	}

	if letstmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not %s . got=%s", name, letstmt.Name)
		return false
	}

	return true
}
