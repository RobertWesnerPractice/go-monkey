package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l            *lexer.Lexer
	tokenCurrent token.Token
	tokenPeek    token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.tokenCurrent = p.tokenPeek
	p.tokenPeek = p.l.NextToken()
}

func (p *Parser) Parse() *ast.Program {
	return nil
}
