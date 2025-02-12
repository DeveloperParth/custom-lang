package ast

import "github.com/developerparth/my-own-lang/tokens"

// statements
type AssignStatement struct {
	Name  tokens.Token
	Value Expression
}

func (a AssignStatement) stmt() {}

type BlockStatement struct {
	Statements []Statement
}

func (b BlockStatement) stmt() {}

type PrintStatement struct {
	Token      tokens.Token
	Expression Expression
}

func (p PrintStatement) stmt() {}

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

type FuncDeclarationStatement struct {
	Identifier tokens.Token
	Block      BlockStatement
	Params     []FunctionParameters
}

func (f FuncDeclarationStatement) stmt() {}

type FunctionParameters struct {
	Token    tokens.Token
	Name     string
	Datatype Type
}

func (f FunctionParameters) stmt() {}

type FuncCallStatement struct {
	Identifier tokens.Token
}

func (f FuncCallStatement) stmt() {}
