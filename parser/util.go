package parser

import "github.com/developerparth/my-own-lang/tokens"

func (p *Parser) match(tokenType ...tokens.Type) bool {
	for _, t := range tokenType {
		if p.token.TokenType == t {
			return true
		}
	}
	return false
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
