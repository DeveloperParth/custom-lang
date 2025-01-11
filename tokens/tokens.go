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
	STRING

	// operators
	ASSIGN
	PLUS
	MINUS
	SLASH
	STAR
	BANG_EQUAL
	EQUAL_EQUAL
	GREATER_THAN
	GREATER_THAN_EQUAL
	LESS_THAN
	LESS_THAN_EQUAL

	// symbols
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE

	// keywords
	PRINT
	TRUE
	FALSE
	NULL

	// special tokens
	EOL
	EOF
)

func (t Type) String() string {
	switch t {
	case ASSIGN:
		return "ASSIGN"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case SLASH:
		return "SLASH"
	case STAR:
		return "STAR"
	case INT:
		return "INT"
	case IDENTIFIER:
		return "IDENTIFIER"
	case STRING:
		return "STRING"
	case EOL:
		return "EOL"
	case EOF:
		return "EOF"
	case PRINT:
		return "PRINT"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	case NULL:
		return "NULL"
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	case LEFT_BRACE:
		return "LEFT_BRACE"
	case RIGHT_BRACE:
		return "RIGHT_BRACE"
	case Illegal:
		return "ILLEGAL"
	default:
		return "UNKNOWN " + string(rune(t))
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
	"true":  TRUE,
	"false": FALSE,
	"null":  NULL,
}
