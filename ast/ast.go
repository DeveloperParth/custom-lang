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

// statements
type AssignStatement struct {
	Token tokens.Token
	Name  *Identifier
	Value Node
}

func (a *AssignStatement) stmt() {}

type BlockStatement struct {
	Statements []Statement
}

func (b *BlockStatement) stmt() {}

type Identifier struct {
	Token tokens.Token
	Value string
}

type PrintStatement struct {
	Token      tokens.Token
	Expression Node
}

func (p *PrintStatement) stmt() {}

type ExpressionStatement struct {
	Expression Expression
}

func (e *ExpressionStatement) stmt() {}
