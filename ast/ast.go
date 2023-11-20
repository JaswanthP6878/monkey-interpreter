package ast

import "monkey/token"

// this is the fundamental node interface
type Node interface {
	TokenLiteral() string
}

// There are two types of nodes to be constructed by the AST
// Statments are code that do not return a value
type Statement interface {
	Node
	statementNode()
}

// Expressions have a return value
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

// let statement

type LetStatement struct {
	Token token.Token // let Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
