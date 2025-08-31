package ast

import (
	"monkey/token"
)

type BooleanExpression struct {
	Token token.Token
	Value bool
}

func (b *BooleanExpression) expressionNode() {}

func (b *BooleanExpression) TokenLiteral() string {
	return b.Token.Literal
}

func (b *BooleanExpression) String() string {
	return b.Token.Literal
}
