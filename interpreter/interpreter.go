package interpreter

import (
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
		println(value)
	case *ast.AssignStatement:
		value := interpretExpression(statement.Value, env)
		env.setInt(statement.Name.Value, value)
	}

}

func interpretBlockStatement(block *ast.BlockStatement, env *Environment) {
	newEnv := NewEnvironment(env)
	for _, statement := range block.Statements {
		interpret(statement, newEnv)
	}
}

func interpretExpression(expression ast.Expression, env *Environment) int64 {
	switch expression := expression.(type) {
	case *ast.IntegerExpr:
		return expression.Value
	case *ast.BinaryExpr:
		left := interpretExpression(expression.Left, env)
		right := interpretExpression(expression.Right, env)

		switch expression.Operator.TokenType {
		case tokens.PLUS:
			return left + right
		case tokens.MINUS:
			return left - right
		case tokens.STAR:
			return left * right
		case tokens.SLASH:
			return left / right
		default:
			panic("Unknown operator")
		}

	case *ast.IdentifierExpr:
		variable := env.getInt(expression.Name)
		return variable
	default:
		panic("Unknown expression")
	}
}
