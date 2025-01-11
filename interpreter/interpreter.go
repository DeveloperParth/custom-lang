package interpreter

import (
	"fmt"

	"github.com/developerparth/my-own-lang/ast"
	"github.com/developerparth/my-own-lang/tokens"
)

func Interpret(statements []ast.Statement) {
	initalEnv := NewEnvironment(nil)
	for _, statement := range statements {
		interpret(statement, initalEnv)
	}
}

func interpret(statement ast.Statement, env *Environment) {
	switch statement := statement.(type) {
	case *ast.BlockStatement:
		interpretBlockStatement(statement, env)
	case *ast.ExpressionStatement:
		interpretExpression(statement.Expression, env)
	case *ast.PrintStatement:
		value := interpretExpression(statement.Expression, env)
		fmt.Println(value.value)
	case *ast.IfStatement:
		condition := interpretExpression(statement.Condition, env)
		if condition.datatype != BOOL {
			panic("Condition must be a boolean")
		}
		if condition.value.(bool) {
			interpret(statement.Then, env)
		} else {
			if statement.Else != nil {
				interpret(statement.Else, env)
			}
		}
	case *ast.AssignStatement:
		value := interpretExpression(statement.Value, env)
		// check if the variable is already defined
		existingVar, exists := env.get(statement.Name.Value)
		if exists {
			if existingVar.datatype != value.datatype {
				message := fmt.Sprintf("Cannot assign %v to %v", value.datatype, existingVar.datatype)
				panic(message)
			}
			existingVar.Literal = value
		} else {
			env.set(statement.Name.Value, value)
		}

	default:
		println(statement)
		panic("Unknown statement")
	}

}

func interpretBlockStatement(block *ast.BlockStatement, env *Environment) {
	newEnv := NewEnvironment(env)
	for _, statement := range block.Statements {
		interpret(statement, newEnv)
	}
}

func interpretExpression(expression ast.Expression, env *Environment) Literal {
	switch expression := expression.(type) {
	case *ast.IntegerExpr:
		return NewLiteral(INT, expression.Value)
	case *ast.BinaryExpr:
		leftLit := interpretExpression(expression.Left, env)
		rightLit := interpretExpression(expression.Right, env)

		if leftLit.datatype != INT || rightLit.datatype != INT {
			message := fmt.Sprintf("Unsupported operation between %v and %v", leftLit.datatype, rightLit.datatype)
			panic(message)
		}
		left := leftLit.value.(int64)
		right := rightLit.value.(int64)
		switch expression.Operator.TokenType {
		case tokens.PLUS:
			return NewLiteral(INT, left+right)
		case tokens.MINUS:
			return NewLiteral(INT, left-right)
		case tokens.STAR:
			return NewLiteral(INT, left*right)
		case tokens.SLASH:
			return NewLiteral(INT, left/right)
		case tokens.GREATER_THAN:
			return NewLiteral(BOOL, left > right)
		case tokens.GREATER_THAN_EQUAL:
			return NewLiteral(BOOL, left >= right)
		case tokens.LESS_THAN:
			return NewLiteral(BOOL, left < right)
		case tokens.LESS_THAN_EQUAL:
			return NewLiteral(BOOL, left <= right)
		case tokens.EQUAL_EQUAL:
			return NewLiteral(BOOL, left == right)
		case tokens.BANG_EQUAL:
			return NewLiteral(BOOL, left != right)
		default:
			panic("Unknown operator")
		}

	case *ast.IdentifierExpr:
		variable := env.getOrPanic(expression.Name)
		return variable.Literal
	case *ast.BooleanExpr:
		return NewLiteral(BOOL, expression.Value)
	default:
		panic("Unknown expression")
	}
}
