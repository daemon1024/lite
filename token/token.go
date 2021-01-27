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

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

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

var keywords = map[string]TokenType{
	"let":   LET,
	"true":  TRUE,
	"false": FALSE,
	"if":    IF,
	"then":  THEN,
	"else":  ELSE,
}

//LookupIdent looks up if the given identifier is a keyword or not
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
