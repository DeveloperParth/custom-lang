package parser

import (
	"github.com/developerparth/my-own-lang/ast"
	"github.com/developerparth/my-own-lang/tokens"
)

func (p *Parser) parseBlockStatement() ast.Statement {
	statements := make([]ast.Statement, 0)
	p.expect(tokens.LEFT_BRACE_CURLY)
	for !p.match(tokens.RIGHT_BRACE_CURLY) {
		statement := p.parse(p.token)
		if statement != nil {
			statements = append(statements, statement)
		}
	}
	p.expect(tokens.RIGHT_BRACE_CURLY)
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

func (p *Parser) parseFuncDeclaration() ast.Statement {
	params := make([]ast.FunctionParameters, 0)
	p.expect(tokens.FUNC)
	identifier := p.expect(tokens.IDENTIFIER)
	p.expect(tokens.LEFT_PAREN)
	// todo: handle params
	for !p.match(tokens.RIGHT_PAREN) {
		// varaible name
		ident := p.expect(tokens.IDENTIFIER)
		// type
		datatype := parse_type(p)
		if !p.match(tokens.RIGHT_PAREN) {
			p.expect(tokens.COMMA)
		}
		params = append(params, ast.FunctionParameters{
			Datatype: datatype,
			Name:     ident.Value,
			Token:    ident,
		})
	}
	p.expect(tokens.RIGHT_PAREN)
	block := p.parseBlockStatement()
	blockStatements, ok := block.(*ast.BlockStatement)
	if !ok {
		panic("Invalid block statement")
	}
	return &ast.FuncDeclarationStatement{
		Identifier: identifier,
		Block:      *blockStatements,
		Params:     params,
	}
}
