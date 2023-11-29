package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestParser(t *testing.T) {
	input := `
	let x =  5;
	let  y = 10;
	let foobar = 838383;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
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

func checkParserErrors(t *testing.T, p *Parser) {
	if len(p.errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(p.errors))
	for _, err := range p.errors {
		t.Errorf("parser error %q", err)
	}
	t.FailNow()
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return add(15);
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	// check for parser errors
	checkParserErrors(t, p)
	if len(program.Statements) != 3 {
		t.Errorf("Expected 3 statements got %d", len(program.Statements))
	}

	for _, statment := range program.Statements {
		returnStatment, ok := statment.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("statement is not a return statement. got %T", returnStatment)
			continue
		}
		if returnStatment.TokenLiteral() != "return" {
			t.Errorf("statment tokenLiteral is not `return`, got %q", returnStatment.TokenLiteral())
		}
	}
}
