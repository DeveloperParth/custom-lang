package tokens

type Type int

type Token struct {
	TokenType Type
	Value     string
}

// token types
//
//go:generate stringer -type=Type
const (
	Illegal Type = iota // used for illegal / unknown token types
	INT
	IDENTIFIER

	// operators
	ASSIGN
	PLUS

	// symbols
	LEFT_PAREN
	RIGHT_PAREN

	// special tokens
	PRINT
	EOL
)

func (t Type) String() string {
	switch t {
	case ASSIGN:
		return "ASSIGN"
	case PLUS:
		return "PLUS"
	case INT:
		return "INT"
	case IDENTIFIER:
		return "IDENTIFIER"
	case EOL:
		return "EOL"
	case PRINT:
		return "PRINT"
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	default:
		return "ILLEGAL"
	}
}

func NewToken(tokenType Type, value string) Token {
	return Token{
		TokenType: tokenType,
		Value:     value,
	}
}

var Keywords = map[string]Type{
	"print": PRINT,
}
