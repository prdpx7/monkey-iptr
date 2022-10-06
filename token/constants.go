package token

const (

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

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

	EQ     = "=="
	NOT_EQ = "!="

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	BANG     = "!"

	ASSIGN = "="

	LT = "<"
	GT = ">"

	// Special Tokens
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"
)

var Keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}
