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

type BooleanExpr struct {
	Token tokens.Token
	Value bool
}

func (b BooleanExpr) expr() {}

type StringExpr struct {
	Token tokens.Token
	Value string
}

func (s StringExpr) expr() {}

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
	Expression Expression
}

func (p *PrintStatement) stmt() {}

type ExpressionStatement struct {
	Expression Expression
}

func (e ExpressionStatement) stmt() {}

type IfStatement struct {
	Condition Expression
	Then      Statement
	Else      Statement
}

func (i IfStatement) stmt() {}
