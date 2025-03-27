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
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.tokenCurrent.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.tokenCurrent.Type {
	case token.Let:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.tokenCurrent}

	if !p.expectPeek(token.Identifier) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.tokenCurrent, Value: p.tokenCurrent.Literal}

	if !p.expectPeek(token.Assignment) {
		return nil
	}

	// TODO: we are skipping the expressions until we encounter a semicolon
	for !p.currentIs(token.Semicolon) {
		p.nextToken()
	}

	return stmt
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

	return false
}
