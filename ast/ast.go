package ast

import (
	"github.com/developerparth/my-own-lang/tokens"
)

type Node interface {
}

type Statement interface {
	stmt()
}

type Expression interface {
	expr()
}

// Expressions
type IntegerExpr struct {
	Token tokens.Token
	Value int64
}

func (i *IntegerExpr) expr() {}

type BinaryExpr struct {
	Left     Expression
	Operator tokens.Token
	Right    Expression
}

func (b *BinaryExpr) expr() {}

type IdentifierExpr struct {
	Token tokens.Token
	Name  string
}

func (i IdentifierExpr) expr() {}

// statements
type AssignStatement struct {
	Name  tokens.Token
	Value Expression
}

func (a *AssignStatement) stmt() {}

type BlockStatement struct {
	Statements []Statement
}

func (b *BlockStatement) stmt() {}

type PrintStatement struct {
	Token      tokens.Token
	Expression Node
}

func (p *PrintStatement) stmt() {}

type ExpressionStatement struct {
	Expression Expression
}

// stmt implements Statement.
func (e ExpressionStatement) stmt() {}
