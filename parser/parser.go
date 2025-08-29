package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l            *lexer.Lexer
	tokenCurrent token.Token
	tokenPeek    token.Token
	errors       []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

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

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.currentIs(token.EOF) {
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
	case token.Return:
		return p.parseReturnStatement()
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

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.tokenCurrent}

	p.nextToken()

	// TODO: We're skipping the expressinos until we envounter a semicolon
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

	p.peekError(t)

	return false
}
