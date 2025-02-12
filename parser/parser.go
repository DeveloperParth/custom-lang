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

func (p *Parser) Parse(input string) *ast.BlockStatement {
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

		statement := p.parse(&token)
		if statement != nil {
			statements = append(statements, statement)
		}
	}

	root := &ast.BlockStatement{
		Statements: statements,
	}
	litter.Dump(root)
	return root
}

func (p *Parser) parse(token *tokens.Token) ast.Statement {
	switch token.TokenType {
	case tokens.PLUS, tokens.MINUS, tokens.STAR, tokens.SLASH, tokens.INT:
		return p.parseExpressionStatement()
	case tokens.IDENTIFIER:
		return p.parseAssignment()
	case tokens.LEFT_BRACE_CURLY:
		return p.parseBlockStatement()
	case tokens.PRINT:
		return p.parsePrintStatement()
	case tokens.FUNC:
		return p.parseFuncDeclaration()
	case tokens.TRUE, tokens.FALSE:
		return p.parseExpressionStatement()
	case tokens.IF:
		return p.parseIfStatement()
	case tokens.EOL:
		p.next()
		return nil

	default:
		fmt.Println(token.TokenType)
		panic("Invalid token")
	}
}

func (p *Parser) parseAssignment() ast.Statement {
	identifier := p.expect(tokens.IDENTIFIER)
	if p.token.TokenType == tokens.ASSIGN {
		p.expect(tokens.ASSIGN)
		expression := p.parseExpression()
		return &ast.AssignStatement{
			Name:  identifier,
			Value: expression,
		}
	}
	if p.token.TokenType == tokens.LEFT_PAREN {
		p.expect(tokens.LEFT_PAREN)
		// todo : handle params
		p.expect(tokens.RIGHT_PAREN)
		return &ast.FuncCallStatement{
			Identifier: identifier,
		}
	}
	message := fmt.Sprintf("Expected ASSIGN or LEFT_PAREN but got %v", p.token.TokenType.String())
	panic(message)
}
