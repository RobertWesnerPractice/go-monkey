package token

type Token struct {
	Type    Type
	Literal string
	Line    int
	Column  int
}

func New(tokenType Type, literal string, line int, column int) Token {
	return Token{tokenType, literal, line, column}
}
