package ast

import "monkey/token"

type NumberLiteral struct {
	Token token.Token
	Value float64
}

func (il *NumberLiteral) expressionNode() {}
func (il *NumberLiteral) TokenLiteral() string {
	return il.Token.Literal
}
func (il *NumberLiteral) String() string {
	return il.Token.Literal
}
