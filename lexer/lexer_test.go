package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10.00;

let add = fn(a, b) {
	a + b;
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
10 >= 1;
10 <= 100;
`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.Declaration, "let"},
		{token.Identifier, "five"},
		{token.Assignment, "="},
		{token.Number, "5"},
		{token.Semicolon, ";"},
		{token.Declaration, "let"},
		{token.Identifier, "ten"},
		{token.Assignment, "="},
		{token.Number, "10.00"},
		{token.Semicolon, ";"},
		{token.Declaration, "let"},
		{token.Identifier, "add"},
		{token.Assignment, "="},
		{token.Function, "fn"},
		{token.ParenthesisLeft, "("},
		{token.Identifier, "a"},
		{token.Comma, ","},
		{token.Identifier, "b"},
		{token.ParenthesisRight, ")"},
		{token.BraceLeft, "{"},
		{token.Identifier, "a"},
		{token.Plus, "+"},
		{token.Identifier, "b"},
		{token.Semicolon, ";"},
		{token.BraceRight, "}"},
		{token.Semicolon, ";"},
		{token.Declaration, "let"},
		{token.Identifier, "result"},
		{token.Assignment, "="},
		{token.Identifier, "add"},
		{token.ParenthesisLeft, "("},
		{token.Identifier, "five"},
		{token.Comma, ","},
		{token.Identifier, "ten"},
		{token.ParenthesisRight, ")"},
		{token.Semicolon, ";"},
		{token.Not, "!"},
		{token.Minus, "-"},
		{token.Division, "/"},
		{token.Multiplication, "*"},
		{token.Number, "5"},
		{token.Semicolon, ";"},
		{token.Number, "5"},
		{token.LessThan, "<"},
		{token.Number, "10"},
		{token.GreaterThan, ">"},
		{token.Number, "5"},
		{token.Semicolon, ";"},
		{token.If, "if"},
		{token.ParenthesisLeft, "("},
		{token.Number, "5"},
		{token.LessThan, "<"},
		{token.Number, "10"},
		{token.ParenthesisRight, ")"},
		{token.BraceLeft, "{"},
		{token.Return, "return"},
		{token.True, "true"},
		{token.Semicolon, ";"},
		{token.BraceRight, "}"},
		{token.Else, "else"},
		{token.BraceLeft, "{"},
		{token.Return, "return"},
		{token.False, "false"},
		{token.Semicolon, ";"},
		{token.BraceRight, "}"},
		{token.Number, "10"},
		{token.Equal, "=="},
		{token.Number, "10"},
		{token.Semicolon, ";"},
		{token.Number, "10"},
		{token.NotEqual, "!="},
		{token.Number, "9"},
		{token.Semicolon, ";"},
		{token.Number, "10"},
		{token.GreaterOrEqual, ">="},
		{token.Number, "1"},
		{token.Semicolon, ";"},
		{token.Number, "10"},
		{token.LessOrEqual, "<="},
		{token.Number, "100"},
		{token.Semicolon, ";"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		nextToken := l.NextToken()
		if nextToken.Type != tt.expectedType {
			t.Fatalf(
				"tests[%d] - wrong tokentype, expected %q, got %q at line %d column %d (%q)",
				i,
				tt.expectedType.Debug(),
				nextToken.Type.Debug(),
				nextToken.Line,
				nextToken.Column,
				nextToken.Literal,
			)
		}

		if nextToken.Literal != tt.expectedLiteral {
			t.Fatalf(
				"tests[%d] - wrong literal, expected %q, got %q at line %d column %d (%q)",
				i,
				tt.expectedLiteral,
				nextToken.Literal,
				nextToken.Line,
				nextToken.Column,
				nextToken.Literal,
			)
		}
	}
}
