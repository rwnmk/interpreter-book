package lexer

import (
	"interpreter-book/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input :=
		`
		let five = 5;
		let ten = 10;
	
		let add = fn(x, y) {
			x + y;
		};
	
		let result = add(five, ten);
		!-/*5;
		5 < 10 > 5;

		if (5 < 10) {
			return true;
		} else {
			return false;
		}

		10 == 10;
		10 != 9;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SCOL, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="}, // TODO: above tokens in order
		{token.INT, "10"},
		{token.SCOL, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FN, "fn"},
		{token.LP, "("},
		{token.IDENT, "x"},
		{token.COM, ","},
		{token.IDENT, "y"},
		{token.RP, ")"},
		{token.LB, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SCOL, ";"},
		{token.RB, "}"},
		{token.SCOL, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LP, "("},
		{token.IDENT, "five"},
		{token.COM, ","},
		{token.IDENT, "ten"},
		{token.RP, ")"},
		{token.SCOL, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.FSLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SCOL, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SCOL, ";"},
		{token.IF, "if"},
		{token.LP, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RP, ")"},
		{token.LB, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SCOL, ";"},
		{token.RB, "}"},
		{token.ELSE, "else"},
		{token.LB, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SCOL, ";"},
		{token.RB, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SCOL, ";"},
		{token.INT, "10"},
		{token.NEQ, "!="},
		{token.INT, "9"},
		{token.SCOL, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
