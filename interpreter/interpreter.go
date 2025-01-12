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
		message := fmt.Sprintf("Unknown statement: %v", statement)
		panic(message)
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
		return interpretBinaryExpr(expression, env)

	case *ast.IdentifierExpr:
		variable := env.getOrPanic(expression.Name)
		return variable.Literal
	case *ast.BooleanExpr:
		return NewLiteral(BOOL, expression.Value)
	case *ast.StringExpr:
		return NewLiteral(STRING, expression.Value)
	default:
		panic("Unknown expression")
	}
}

func interpretBinaryExpr(expression *ast.BinaryExpr, env *Environment) Literal {
	leftLit := interpretExpression(expression.Left, env)
	rightLit := interpretExpression(expression.Right, env)

	arethmeticOperators := []tokens.Type{
		tokens.PLUS,
		tokens.MINUS,
		tokens.STAR,
		tokens.SLASH,
	}
	for _, operator := range arethmeticOperators {
		if expression.Operator.TokenType == operator {
			return interpretArithmeticBinaryExpr(expression, leftLit, rightLit)
		}
	}
	if leftLit.datatype != rightLit.datatype {
		message := fmt.Sprintf("Unsupported %v operation between %v and %v", expression.Operator.TokenType.String(), leftLit.datatype, rightLit.datatype)
		panic(message)
	}
	if leftLit.datatype == INT {
		left := leftLit.value.(int64)
		right := rightLit.value.(int64)
		switch expression.Operator.TokenType {
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
	} else if leftLit.datatype == BOOL {
		left := leftLit.value.(bool)
		right := rightLit.value.(bool)
		switch expression.Operator.TokenType {
		case tokens.EQUAL_EQUAL:
			return NewLiteral(BOOL, left == right)
		case tokens.BANG_EQUAL:
			return NewLiteral(BOOL, left != right)
		default:
			panic("Unknown operator")
		}
	} else {
		panic("Unknown datatype")
	}
}

func interpretArithmeticBinaryExpr(expression *ast.BinaryExpr, leftLit, rightLit Literal) Literal {
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
	default:
		panic("Unknown operator")
	}
}
