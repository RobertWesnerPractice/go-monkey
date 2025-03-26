package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.Assignment, "="},
		{token.Plus, "+"},
		{token.ParenthesisLeft, "("},
		{token.ParenthesisRight, ")"},
		{token.BraceLeft, "{"},
		{token.BraceRight, "}"},
		{token.Comma, ","},
		{token.Semicolon, ";"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		nextToken := l.NextToken()
		if nextToken.Type != tt.expectedType {
			t.Fatalf(
				"tests[%d] - wrong tokentype, expected %q, got %q at column %d (%q)",
				i,
				tt.expectedType.Debug(),
				nextToken.Type.Debug(),
				nextToken.Column,
				nextToken.Literal,
			)
		}

		if nextToken.Literal != tt.expectedLiteral {
			t.Fatalf(
				"tests[%d] - wrong literal, expected %q, got %q at column %d",
				i,
				tt.expectedLiteral,
				nextToken.Literal,
				nextToken.Column,
			)
		}
	}
}
