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
	var tok token.Token

	l.eatWhitespace()

	switch l.ch {
	case '=':
		tok = l.createToken(token.Assignment, string(l.ch))
	case '+':
		tok = l.createToken(token.Plus, string(l.ch))
	case ',':
		tok = l.createToken(token.Comma, string(l.ch))
	case ';':
		tok = l.createToken(token.Semicolon, string(l.ch))
	case '(':
		tok = l.createToken(token.ParenthesisLeft, string(l.ch))
	case ')':
		tok = l.createToken(token.ParenthesisRight, string(l.ch))
	case '{':
		tok = l.createToken(token.BraceLeft, string(l.ch))
	case '}':
		tok = l.createToken(token.BraceRight, string(l.ch))
	case 0:
		tok = l.createToken(token.EOF, "")
	default:
		if isLetter(l.ch) {
			identifier := l.readIdentifier()
			return l.createToken(token.LookupIdentifier(identifier), identifier)
		}

		if isDigit(l.ch) {
			return l.createToken(token.Number, l.readNumber())
		}

		return l.createToken(token.Illegal, string(l.ch))
	}

	l.readChar()

	return tok
}

func (l *Lexer) createToken(tokenType token.Type, literal string) token.Token {
	return token.New(tokenType, literal, l.line, l.column)
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	// float support
	if l.ch == '.' {
		l.readChar()
		for isDigit(l.ch) {
			l.readChar()
		}
	}

	return l.input[position:l.position]
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
