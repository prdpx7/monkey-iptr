package token

const (
	// Keywords
	FUNCTION = "FUCTION"
	LET      = "LET"

	// Identifiers
	INT    = "INT"
	IDENT  = "IDENT"
	STRING = "STRING"
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN         = "("
	RPAREN         = ")"
	LBRACE         = "{"
	RBRACE         = "}"
	SLASH          = "/"
	BACKWARD_SLASH = "\\"

	// Operators
	PLUS       = "+"
	MINUS      = "-"
	STAR       = "*"
	MATH_POWER = "**"
	ASSIGN     = "="

	// Special Tokens
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"
)

var Keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}
