package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/token"
	"strconv"
)

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

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.tokenCurrent,
		Operator: p.tokenCurrent.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(Prefix)

	return expression
}
