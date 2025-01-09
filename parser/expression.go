package parser

import (
	"fmt"
	"strconv"

	"github.com/developerparth/my-own-lang/ast"
	"github.com/developerparth/my-own-lang/tokens"
)

func (p *Parser) parseExpression() *ast.ExpressionStatement {
	return &ast.ExpressionStatement{
		Expression: p.parseEquality(),
	}
}

func (p *Parser) parseEquality() ast.Expression {
	left := p.parseComparison()
	tokensToMatch := []tokens.Type{
		tokens.EQUAL_EQUAL,
		tokens.BANG_EQUAL,
	}
	for p.match(tokensToMatch...) {
		operator := p.expect(tokensToMatch...)
		right := p.parseComparison()
		left = &ast.BinaryExpr{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}
	return left
}

func (p *Parser) parseComparison() ast.Expression {
	left := p.parseTerm()
	tokensToMatch := []tokens.Type{
		tokens.GREATER_THAN,
		tokens.GREATER_THAN_EQUAL,
		tokens.LESS_THAN,
		tokens.LESS_THAN_EQUAL,
	}
	for p.match(tokensToMatch...) {
		operator := p.expect(tokensToMatch...)
		right := p.parseTerm()
		left = &ast.BinaryExpr{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}
	return left
}

func (p *Parser) parseTerm() ast.Expression {
	left := p.parseFactor()
	toknesToMatch := []tokens.Type{
		tokens.PLUS,
		tokens.MINUS,
	}
	for p.match(toknesToMatch...) {
		operator := p.expect(toknesToMatch...)
		right := p.parseFactor()
		left = &ast.BinaryExpr{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}
	return left
}

func (p *Parser) parseFactor() ast.Expression {
	left := p.parsePrimary()
	tokensToMatch := []tokens.Type{
		tokens.STAR,
		tokens.SLASH,
	}
	for p.match(tokensToMatch...) {
		operator := p.expect(tokensToMatch...)
		right := p.parsePrimary()
		left = &ast.BinaryExpr{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}
	return left

}
func (p *Parser) parsePrimary() ast.Expression {
	if p.match(tokens.INT) {
		token := p.next()
		value, _ := strconv.ParseInt(token.Value, 10, 64)
		return &ast.IntegerExpr{
			Token: token,
			Value: value,
		}
	}
	if p.match(tokens.LEFT_PAREN) {
		p.expect(tokens.LEFT_PAREN)
		expr := p.parseEquality()
		p.expect(tokens.RIGHT_PAREN)
		return expr
	}
	fmt.Println(p.token)
	panic("Invalid token")
}
