package parser

import (
	"unicode"

	"github.com/developerparth/my-own-lang/tokens"
)

type Lexer struct {
	// source code
	input string
	// current position in the input (points to current char)
	position int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:    input,
		position: 0,
	}
}

func (lexer *Lexer) Next() tokens.Token {
	return lexer.next()
}

func (lexer *Lexer) next() tokens.Token {
	if lexer.isEnd() {
		return tokens.NewToken(tokens.EOF, "")
	}
	switch lexer.current() {
	case '+':
		lexer.advance()
		return tokens.NewToken(tokens.PLUS, "+")
	case '-':
		lexer.advance()
		return tokens.NewToken(tokens.MINUS, "-")
	case '*':
		lexer.advance()
		return tokens.NewToken(tokens.STAR, "*")
	case '/':
		lexer.advance()
		return tokens.NewToken(tokens.SLASH, "/")
	case '=':
		lexer.advance()
		return tokens.NewToken(tokens.ASSIGN, "=")
	case '!':
		if lexer.current() == '=' {
			lexer.advance()
			lexer.advance()
			return tokens.NewToken(tokens.BANG_EQUAL, "!=")
		}
		return tokens.NewToken(tokens.Illegal, "")
	case '>':
		lexer.advance()
		if lexer.current() == '=' {
			lexer.advance()
			return tokens.NewToken(tokens.GREATER_THAN_EQUAL, ">=")
		}
		return tokens.NewToken(tokens.GREATER_THAN, ">")
	case '<':
		lexer.advance()
		if lexer.current() == '=' {
			lexer.advance()
			return tokens.NewToken(tokens.LESS_THAN_EQUAL, "<=")
		}
		return tokens.NewToken(tokens.LESS_THAN, "<")
	case '(':
		lexer.advance()
		return tokens.NewToken(tokens.LEFT_PAREN, "(")
	case ')':
		lexer.advance()
		return tokens.NewToken(tokens.RIGHT_PAREN, ")")
	case ';':
		lexer.advance()
		return lexer.next()
	case ' ':
		lexer.advance()
		return lexer.next()
	case '\n':
		lexer.advance()
		return tokens.NewToken(tokens.EOL, "\n")
	default:
		if unicode.IsDigit(rune(lexer.current())) {
			var value string
			for unicode.IsDigit(rune(lexer.current())) {
				char := lexer.advance()
				value += string(char)
				if lexer.isEnd() {
					break
				}
			}
			return tokens.NewToken(tokens.INT, value)
		}
		if unicode.IsLetter(rune(lexer.current())) {
			var value string

			for unicode.IsLetter(rune(lexer.current())) || unicode.IsDigit(rune(lexer.current())) {
				char := lexer.advance()
				value += string(char)
				if lexer.isEnd() {
					break
				}
			}
			return tokens.NewToken(tokens.IDENTIFIER, value)
		}
		return tokens.NewToken(tokens.Illegal, "")

	}

}

func (lexer *Lexer) current() byte {
	return lexer.input[lexer.position]
}

func (lexer *Lexer) peek() byte {
	return lexer.input[lexer.position+1]
}

func (lexer *Lexer) advance() byte {
	char := lexer.current()
	if !lexer.isEnd() {
		lexer.position++
	}
	return char
}
func (lexer *Lexer) isEnd() bool {
	return lexer.position >= len(lexer.input)
}
