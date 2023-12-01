package ast

import (
	"bytes"
	"monkey/token"
)

// this is the fundamental node interface
type Node interface {
	TokenLiteral() string
	String() string
}

// There are two types of nodes to be constructed by the AST
// Statments are code that do not return a value

// statementNode method is a marker interface method so as to make
// compiler care for statementNode
type Statement interface {
	Node
	statementNode()
}

// Expressions have a return value
// expressionNode method is a marker interface method so as to make
// compiler care for statementNode
type Expression interface {
	Node
	expressionNode()
}

// This is the root node of the AST
// It consists of a slice of statements
type Program struct {
	Statements []Statement
}

// utils programs
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, statement := range p.Statements {
		out.WriteString(statement.String())
	}

	return out.String()
}

// let statement

type LetStatement struct {
	Token token.Token // let Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Identifier is considered expression
// this is for convinence as some Identifies do produce values
type Identifier struct {
	Token token.Token // IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string { return i.Value }

// Return statement has the syntax of
// return <expression>;
type ReturnStatement struct {
	Token       token.Token // RETURN token
	ReturnValue Expression
}

func (r *ReturnStatement) statementNode()       {}
func (r *ReturnStatement) TokenLiteral() string { return r.Token.Literal }

func (r *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(r.TokenLiteral() + " ")

	if r.ReturnValue != nil {
		out.WriteString(r.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// This is a wrapper type around expressions to have them as Statements in
// the ast
type ExpressionStatement struct {
	Token      token.Token // the first token of this expression
	Expression Expression
}

func (r *ExpressionStatement) statementNode()       {}
func (r *ExpressionStatement) TokenLiteral() string { return r.Token.Literal }

func (e *ExpressionStatement) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}

	return ""

}
