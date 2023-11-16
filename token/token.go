package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// indentifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// operators
	ASSIGN  = "="
	PLUS    = "+"
	LT      = "<"
	GT      = ">"
	BANG    = "!"
	ASTERIK = "*"
	SLASH   = "/"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"else":   ELSE,
	"if":     IF,
	"return": RETURN,
}

func LookUpIndent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
