package lexer

import "interpreter-book/token"

// pos/readPos are used to access chars in "input" by using them as an index i.e. l.input[l.readPos]

type Lexer struct {
	input   string
	pos     int  // changed from position :: current position in input (points to current char)
	readPos int  // changes from readPosition :: current reading position in input (after current char)
	ch      byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// gives us the next char and advance our pos in the input string, if we're at the end of input then either
// we've not read anything yet or we're at EOF. The rest is self-explanatory.
// TODO: add unicode support (l.ch -> rune (or []rune?), then instead of reading next byte we need to read next character delimiter)
func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SCOL, l.ch)
	case '(':
		tok = newToken(token.LP, l.ch)
	case ')':
		tok = newToken(token.RP, l.ch)
	case ',':
		tok = newToken(token.COM, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LB, l.ch)
	case '}':
		tok = newToken(token.RB, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
