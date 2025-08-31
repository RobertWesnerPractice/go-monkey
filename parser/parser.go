package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l              *lexer.Lexer
	tokenCurrent   token.Token
	tokenPeek      token.Token
	errors         []string
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

var precedences = map[token.Type]int{
	token.Equal:          Equals,
	token.NotEqual:       Equals,
	token.LessThan:       LessGreater,
	token.GreaterThan:    LessGreater,
	token.Plus:           Sum,
	token.Minus:          Sum,
	token.Multiplication: Product,
	token.Division:       Product,
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()

	p.prefixParseFns = map[token.Type]prefixParseFn{}
	p.registerPrefix(token.Identifier, p.parseIdentifier)
	p.registerPrefix(token.Number, p.parseNumberLiteral)
	p.registerPrefix(token.Not, p.parsePrefixExpression)
	p.registerPrefix(token.Minus, p.parsePrefixExpression)
	p.registerPrefix(token.True, p.parseBooleanExpression)
	p.registerPrefix(token.False, p.parseBooleanExpression)
	p.registerPrefix(token.ParenthesisLeft, p.parseGroupedExpression)

	p.infixParseFns = map[token.Type]infixParseFn{}
	p.registerInfix(token.Plus, p.parseInfixExpression)
	p.registerInfix(token.Minus, p.parseInfixExpression)
	p.registerInfix(token.Multiplication, p.parseInfixExpression)
	p.registerInfix(token.Division, p.parseInfixExpression)
	p.registerInfix(token.Equal, p.parseInfixExpression)
	p.registerInfix(token.NotEqual, p.parseInfixExpression)
	p.registerInfix(token.GreaterThan, p.parseInfixExpression)
	p.registerInfix(token.GreaterOrEqual, p.parseInfixExpression)
	p.registerInfix(token.LessThan, p.parseInfixExpression)
	p.registerInfix(token.LessOrEqual, p.parseInfixExpression)

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
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
