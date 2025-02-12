package interpreter

import "github.com/developerparth/my-own-lang/ast"

type Datatype string

const (
	INT    Datatype = "int"
	BOOL   Datatype = "bool"
	NULL   Datatype = "null"
	STRING Datatype = "string"
)

type Literal struct {
	datatype Datatype
	value    any
}

func NewLiteral(datatype Datatype, value any) Literal {
	return Literal{
		datatype: datatype,
		value:    value,
	}
}

type Variable struct {
	Name string
	Literal
}

func (v *Variable) get() any {
	switch v.Literal.datatype {
	case INT:
		return v.Literal.value.(int64)
	case BOOL:
		return v.Literal.value.(bool)
	case NULL:
		return nil
	case STRING:
		return v.Literal.value.(string)
	}
	return nil
}

type Environment struct {
	parent    *Environment
	vars      map[string]Variable
	functions map[string]Function
}

type Function struct {
	env   *Environment
	block *ast.BlockStatement
}

func (e *Environment) getOrPanic(name string) Variable {
	variable, ok := e.get(name)
	if !ok {
		panic("Unknown variable")
	}
	return variable
}

func (e *Environment) get(name string) (Variable, bool) {
	value, ok := e.vars[name]
	if !ok {
		if e.parent != nil {
			return e.parent.get(name)
		} else {
			return Variable{}, false
		}
	}
	return value, true
}

func (e *Environment) set(name string, value Literal) {
	e.vars[name] = Variable{
		Name:    name,
		Literal: value,
	}
}

func (e *Environment) getFunction(name string) (Function, bool) {
	value, ok := e.functions[name]
	if !ok {
		if e.parent != nil {
			return e.parent.getFunction(name)
		} else {
			return Function{}, false
		}
	}
	return value, true
}

func (e *Environment) setFunction(name string, block *ast.BlockStatement) {
	if _, ok := e.getFunction(name); ok {
		panic("Function already exists")
	}
	e.functions[name] = Function{
		env:   NewEnvironment(e),
		block: block,
	}
}

func NewEnvironment(parent *Environment) *Environment {
	return &Environment{
		vars:      make(map[string]Variable),
		functions: make(map[string]Function),
		parent:    parent,
	}
}
