package parser

import (
	"github.com/developerparth/my-own-lang/ast"
	"github.com/developerparth/my-own-lang/tokens"
)

func parse_type(p *Parser) ast.Type {
	if p.token.TokenType != tokens.IDENTIFIER {
		panic("Not implemented")
	}
	if p.token.Value == "int" {
		p.next()
		return ast.NumberType{}
	}
	panic("Invalid type")

}
