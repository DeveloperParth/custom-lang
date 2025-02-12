package parser

import (
	"fmt"
	"strconv"

	"github.com/developerparth/my-own-lang/ast"
	"github.com/developerparth/my-own-lang/tokens"
)

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	return &ast.ExpressionStatement{
		Expression: p.parseExpression(),
	}
}

func (p *Parser) parseExpression() ast.Expression {
	return p.parseEquality()
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
	if p.match(tokens.IDENTIFIER) {
		token := p.next()
		return &ast.IdentifierExpr{
			Token: token,
			Name:  token.Value,
		}
	}
	if p.match(tokens.LEFT_PAREN) {
		p.expect(tokens.LEFT_PAREN)
		expr := p.parseEquality()
		p.expect(tokens.RIGHT_PAREN)
		return expr
	}
	if p.match(tokens.TRUE, tokens.FALSE) {
		token := p.next()
		value := token.TokenType == tokens.TRUE
		return &ast.BooleanExpr{
			Token: token,
			Value: value,
		}
	}
	if p.match(tokens.STRING) {
		token := p.next()
		return &ast.StringExpr{
			Token: token,
			Value: token.Value,
		}
	}
	if p.match(tokens.LEFT_BRACE) {
		p.expect(tokens.LEFT_BRACE)

		arrayOfElements := make([]ast.Expression, 0)

		for p.token.TokenType != tokens.RIGHT_BRACE {
			arrayOfElements = append(arrayOfElements, p.parseExpression())
			println(p.token.TokenType.String())
			if p.token.TokenType != tokens.RIGHT_BRACE {
				p.expect(tokens.COMMA)
			}
		}
		p.expect(tokens.RIGHT_BRACE)
		return &ast.ArrayExpr{
			Elements: arrayOfElements,
		}
	}
	message := fmt.Sprintf("Unexpected token %v", p.token.TokenType.String())
	panic(message)
}
