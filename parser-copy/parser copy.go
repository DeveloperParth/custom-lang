package parsercopy

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/developerparth/my-own-lang/ast"
	"github.com/developerparth/my-own-lang/tokens"
	"github.com/sanity-io/litter"
)

func appendCurrentToken(currentToken *string, tokensArray *[]tokens.Token) {
	if *currentToken != "" {
		if kind, exists := tokens.Keywords[*currentToken]; exists {
			fmt.Println("Keyword: ", *currentToken)
			token := &tokens.Token{
				Value:     *currentToken,
				TokenType: kind,
			}
			*tokensArray = append(*tokensArray, *token)
			*currentToken = ""
			return
		}
		isInt := true

		for _, char := range *currentToken {
			if !unicode.IsDigit(char) {
				isInt = false
				break
			}
		}
		var token *tokens.Token
		if isInt {
			token = &tokens.Token{
				Value:     *currentToken,
				TokenType: tokens.INT,
			}
		} else {
			token = &tokens.Token{
				Value:     *currentToken,
				TokenType: tokens.IDENTIFIER,
			}
		}
		*tokensArray = append(*tokensArray, *token)
	}
	*currentToken = ""
}
func getTokens(line string) []tokens.Token {

	tokensArray := []tokens.Token{}
	if len(line) == 0 {
		return tokensArray
	}

	i := 0
	char := line[i]
	currentToken := ""

	for i < len(line) {
		if char == ' ' {
			appendCurrentToken(&currentToken, &tokensArray)
		} else if char == '=' {
			token := &tokens.Token{
				Value:     "=",
				TokenType: tokens.ASSIGN,
			}
			appendCurrentToken(&currentToken, &tokensArray)
			tokensArray = append(tokensArray, *token)
		} else if char == '+' {
			token := &tokens.Token{
				Value:     "+",
				TokenType: tokens.PLUS,
			}
			appendCurrentToken(&currentToken, &tokensArray)
			tokensArray = append(tokensArray, *token)
		} else if char == '(' {
			appendCurrentToken(&currentToken, &tokensArray)
			tokensArray = append(tokensArray, tokens.NewToken(tokens.LEFT_PAREN, "("))
		} else if char == ')' {
			appendCurrentToken(&currentToken, &tokensArray)
			tokensArray = append(tokensArray, tokens.NewToken(tokens.RIGHT_PAREN, ")"))
		} else {
			currentToken += string(char)
		}

		if i == len(line)-1 {
			appendCurrentToken(&currentToken, &tokensArray)
			token := &tokens.Token{
				Value:     "",
				TokenType: tokens.EOL,
			}
			tokensArray = append(tokensArray, *token)
			break
		}
		i++
		char = line[i]
	}

	return tokensArray

}

type Parser struct {
	Lines  []string
	token  *tokens.Token
	tokens []tokens.Token
	index  int
}

func (p *Parser) Parse() {
	statements := make([]ast.Statement, 0)
	for _, line := range p.Lines {
		p.tokens = getTokens(line)
		p.index = 0
		p.token = &p.tokens[p.index]
		tree := p.parse(p.token)

		statements = append(statements, tree)

		if tree != nil {
		}
	}
	root := &ast.BlockStatement{
		Statements: statements,
	}
	litter.Dump(root)
}

func expectToken(token *tokens.Token, tokenType ...tokens.Type) {
	for _, t := range tokenType {
		if token.TokenType == t {
			return
		}
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

func (p *Parser) expect(tokenType ...tokens.Type) {
	expectToken(p.token, tokenType...)
}

func (p *Parser) next() tokens.Token {
	token := p.tokens[p.index]
	p.index++
	p.token = &p.tokens[p.index]
	return token
}

func (p *Parser) parse(token *tokens.Token) ast.Statement {
	switch token.TokenType {
	// case tokens.ASSIGN:
	// 	return p.parseAssignment()
	case tokens.PRINT:
		return p.parsePrint()
	case tokens.EOL:
		return nil
	case tokens.INT:
		return p.parseExpression()
	case tokens.IDENTIFIER:
		return p.parseIdentifier()
	default:
		fmt.Println(token.TokenType)
		panic("Invalid token")
	}

}

func (p *Parser) parseIdentifier() *ast.IdentifierStatement {
	p.expect(tokens.IDENTIFIER)
	identifierToken := p.next()
	p.expect(tokens.ASSIGN)
	p.next()
	p.expect(tokens.INT)
	valToken := p.next()
	p.expect(tokens.EOL)

	Value, _ := strconv.ParseInt(valToken.Value, 10, 64)

	return &ast.IdentifierStatement{
		Token: *p.token,
		Name:  identifierToken.Value,
		Value: &ast.IntegerExpr{
			Token: valToken,
			Value: Value,
		},
	}
}

func (p *Parser) parseExpression() *ast.ExpressionStatement {
	var expression ast.Expression
	int, _ := strconv.ParseInt(p.token.Value, 10, 64)

	expression = &ast.IntegerExpr{
		Token: *p.token,
		// cast to int64
		Value: int,
	}

	return &ast.ExpressionStatement{
		Expression: expression,
	}
}

func (p *Parser) parsePrint() *ast.PrintStatement {
	p.expect(tokens.PRINT)
	p.next()
	p.expect(tokens.LEFT_PAREN)
	p.next()
	p.expect(tokens.INT, tokens.IDENTIFIER)
	valueToken := p.next()
	p.expect(tokens.RIGHT_PAREN)
	p.next()
	p.expect(tokens.EOL)

	var value = p.parse(&valueToken)

	node := &ast.PrintStatement{
		Token:      *p.token,
		Expression: value,
	}

	return node
}

func (p *Parser) parseAssignment() *ast.AssignStatement {
	// p.expect(tokens.IDENTIFIER)
	expectToken(&p.tokens[p.index-1], tokens.IDENTIFIER)
	p.expect(tokens.ASSIGN)
	p.next()
	p.expect(tokens.INT, tokens.IDENTIFIER)
	p.next()
	p.expect(tokens.EOL)

	valueToken := p.tokens[p.index-1]
	var value ast.Node = p.parse(&valueToken)

	node := &ast.AssignStatement{
		Token: *p.token,
		Name: &ast.IdentifierStatement{
			Token: *p.token,
			Value: &ast.IntegerExpr{
				Token: p.tokens[p.index-3],
				Value: 0,
			},
		},
		Value: value,
	}

	return node
}
