package parser

import (
	"fmt"

	"github.com/developerparth/my-own-lang/ast"
	"github.com/developerparth/my-own-lang/tokens"
	"github.com/sanity-io/litter"
)

type Parser struct {
	input  string
	token  *tokens.Token
	tokens []tokens.Token
	index  int
}

func (p *Parser) Parse(input string) {
	statements := make([]ast.Statement, 0)

	p.input = input
	lexer := Lexer{
		input: p.input,
	}
	token := lexer.next()

	for {
		if token.TokenType == tokens.EOF {
			p.tokens = append(p.tokens, token)
			break
		}
		p.tokens = append(p.tokens, token)
		token = lexer.next()

	}

	p.token = &p.tokens[p.index]

	for index, token := range p.tokens {
		if index < p.index {
			continue
		}
		if token.TokenType == tokens.EOF {
			break
		}

		fmt.Printf("Token: %v\n", token)

		statement := p.parse(&token)
		if statement != nil {
			statements = append(statements, statement)
		}
	}

	root := &ast.BlockStatement{
		Statements: statements,
	}
	litter.Dump(root)
}

func (p *Parser) expect(tokenType ...tokens.Type) tokens.Token {
	token := p.token

	if p.match(tokenType...) {
		p.next()
		return *token
	}
	if len(tokenType) == 1 {
		panic("Expected " + tokenType[0].String() + " but got " + token.TokenType.String())
	}
	var expected string
	for i, t := range tokenType {
		if i == 0 {
			expected += t.String()
		} else {
			expected += " or " + t.String()
		}
	}
	panic("Expected " + expected + " but got " + token.TokenType.String())
}

func (p *Parser) next() tokens.Token {
	token := p.tokens[p.index]
	p.index++
	p.token = &p.tokens[p.index]
	return token
}

func (p *Parser) parse(token *tokens.Token) ast.Statement {
	switch token.TokenType {
	case tokens.PLUS, tokens.MINUS, tokens.STAR, tokens.SLASH, tokens.INT:
		fmt.Printf("Parsing expression: %v\n", token)
		return p.parseExpression()
	default:
		fmt.Println(token.TokenType)
		panic("Invalid token")
	}
}

func (p *Parser) match(tokenType ...tokens.Type) bool {
	for _, t := range tokenType {
		if p.token.TokenType == t {
			return true
		}
	}
	return false
}
