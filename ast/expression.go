package ast

import "github.com/developerparth/my-own-lang/tokens"

// Expressions
type IntegerExpr struct {
	Token tokens.Token
	Value int64
}

func (i IntegerExpr) expr() {}

type BinaryExpr struct {
	Left     Expression
	Operator tokens.Token
	Right    Expression
}

func (b BinaryExpr) expr() {}

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

type ArrayExpr struct {
	Elements []Expression
}

func (n ArrayExpr) expr() {}
