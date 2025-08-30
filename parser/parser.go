package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"strconv"
)

type Parser struct {
	l            *lexer.Lexer
	tokenCurrent token.Token
	tokenPeek    token.Token
	errors       []string

	prefixParseFns map[token.Type]prefixParseFn
	infixParseFns  map[token.Type]infixParseFn
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

const (
	_ int = iota
	Lowest
	Equals
	LessGreater
	Sum
	Product
	Prefix
	Call
)

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()

	p.prefixParseFns = map[token.Type]prefixParseFn{}
	p.registerPrefix(token.Identifier, p.parseIdentifier)
	p.registerPrefix(token.Number, p.parseNumberLiteral)
	p.registerPrefix(token.Not, p.parsePrefixExpression)
	p.registerPrefix(token.Minus, p.parsePrefixExpression)

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
		return p.parseExpressionStatement()
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

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.tokenCurrent}

	stmt.Expression = p.parseExpression(Lowest)

	if p.peekIs(token.Semicolon) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.tokenCurrent.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.tokenCurrent.Type)

		return nil
	}
	leftExp := prefix()

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.tokenCurrent, Value: p.tokenCurrent.Literal}
}

func (p *Parser) parseNumberLiteral() ast.Expression {
	lit := &ast.NumberLiteral{Token: p.tokenCurrent}

	value, err := strconv.ParseFloat(p.tokenCurrent.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as number", p.tokenCurrent.Literal)
		p.errors = append(p.errors, msg)

		return nil
	}

	lit.Value = value

	return lit
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

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.tokenCurrent,
		Operator: p.tokenCurrent.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(Prefix)

	return expression
}
