package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Keywords
	LET   = "LET"
	TRUE  = "TRUE"
	FALSE = "FALSE"
	IF    = "IF"
	THEN  = "THEN"
	ELSE  = "ELSE"
)

//TokenType is the type of token
type TokenType string

//Token is any identifier with a meaning
type Token struct {
	Type    TokenType
	Literal string
}
