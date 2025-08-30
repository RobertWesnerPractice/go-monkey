package parser

import (
	"fmt"
	"monkey/token"
)

func (p *Parser) peekError(t token.Type) {
	p.errors = append(
		p.errors,
		fmt.Sprintf("expected next token to be %v, got %v instead", t.Debug(), p.tokenPeek.Type.Debug()),
	)
}

func (p *Parser) nextToken() {
	p.tokenCurrent = p.tokenPeek
	p.tokenPeek = p.l.NextToken()
}

func (p *Parser) currentIs(t token.Type) bool {
	return p.tokenCurrent.Type == t
}

func (p *Parser) peekIs(t token.Type) bool {
	return p.tokenPeek.Type == t
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekIs(t) {
		p.nextToken()

		return true
	}

	p.peekError(t)

	return false
}

func (p *Parser) registerPrefix(tokenType token.Type, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.Type, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) noPrefixParseFnError(t token.Type) {
	msg := fmt.Sprintf("no prefix parse function for %v found", t.Debug())
	p.errors = append(p.errors, msg)
}

func (p *Parser) currentPrecedence() int {
	if p, ok := precedences[p.tokenCurrent.Type]; ok {
		return p
	}

	return Lowest
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.tokenPeek.Type]; ok {
		return p
	}

	return Lowest
}
