package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	line         int
	column       int
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.line, l.column = 1, 1
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++

	if l.ch == 0xA {
		l.column = 0
		l.line++
	} else if l.ch != 0xD {
		l.column++
	}
}

func (l *Lexer) NextToken() token.Token {
	defer l.readChar()

	switch l.ch {
	case '=':
		return l.createToken(token.Assignment, l.ch)
	case '+':
		return l.createToken(token.Plus, l.ch)
	case ',':
		return l.createToken(token.Comma, l.ch)
	case ';':
		return l.createToken(token.Semicolon, l.ch)
	case '(':
		return l.createToken(token.ParenthesisLeft, l.ch)
	case ')':
		return l.createToken(token.ParenthesisRight, l.ch)
	case '{':
		return l.createToken(token.BraceLeft, l.ch)
	case '}':
		return l.createToken(token.BraceRight, l.ch)
	case 0:
		t := l.createToken(token.EOF, l.ch)
		t.Literal = ""

		return t
	default:
		return l.createToken(token.Illegal, l.ch)
	}
}

func (l *Lexer) createToken(tokenType token.Type, ch byte) token.Token {
	return token.New(tokenType, string(ch), l.line, l.column)
}
