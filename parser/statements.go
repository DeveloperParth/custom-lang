package parser

import (
	"github.com/developerparth/my-own-lang/ast"
	"github.com/developerparth/my-own-lang/tokens"
)

func (p *Parser) parseBlockStatement() ast.Statement {
	statements := make([]ast.Statement, 0)
	p.expect(tokens.LEFT_BRACE)
	for !p.match(tokens.RIGHT_BRACE) {
		statement := p.parse(p.token)
		if statement != nil {
			statements = append(statements, statement)
		}
	}
	p.expect(tokens.RIGHT_BRACE)
	return &ast.BlockStatement{
		Statements: statements,
	}
}

func (p *Parser) parsePrintStatement() ast.Statement {
	p.expect(tokens.PRINT)
	expression := p.parseExpression()
	p.expect(tokens.EOL)
	return &ast.PrintStatement{
		Expression: expression,
	}
}

func (p *Parser) parseIfStatement() ast.Statement {
	p.expect(tokens.IF)
	condition := p.parseExpression()
	block := p.parseBlockStatement()
	var elseBlock ast.Statement
	if p.match(tokens.ELSE) {
		p.expect(tokens.ELSE)
		if p.match(tokens.IF) {
			elseBlock = p.parseIfStatement()
		} else {
			elseBlock = p.parseBlockStatement()
		}
	}

	return &ast.IfStatement{
		Condition: condition,
		Then:      block,
		Else:      elseBlock,
	}
}
