// define the token datastructure
// the type attribute takes in a string and we derive the type from the character
// 		> the use of string allows use to use many different values as TokenTypes
// the Literal attribute holds the literal value of the token for reuse and to avoid loss

package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// note to self: changed some token consts to make them shorter (and so marginally more efficient...? investigate)
const (
	// special types
	ILLEGAL = "ILLEGAL" // token/char we don't know about
	EOF     = "EOF"     // end of file (stop parser)

	// identifiers
	IDENT = "IDENT" // add, foobar, x, y etc..
	INT   = "INT"   // 123456 etc

	// operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	FSLASH   = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NEQ      = "!=" // NOT_EQ

	// delimiters
	COM  = "," // COMMA
	SCOL = ";" // SEMICOLON

	LP = "(" // LPAREN
	RP = ")" // RPAREN
	LB = "{" // LBRACE
	RB = "}" // RBRACE

	// keywords
	FN     = "FN" // FUNCTION
	LET    = "LET"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
)

// distinguish user-defined identifiers from language keywords
var keywords = map[string]TokenType{
	"fn":     FN,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
