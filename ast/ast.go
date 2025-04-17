package ast

import (
	token "github.com/TusharAbhinav/monkey/token"
)

type Node interface {
	TokenLiteral() string
}

// Statement and Expression have been inherited from Node interface
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

// Root node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement implements Statement Interface
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier //implements Expression interface
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type ReturnStatement struct {
	Token token.Token // the token.RETURN token
	ReturnValue Expression
}
func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// Identifier implements Expression interface
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
