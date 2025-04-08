package lexer

import (
	token "github.com/TusharAbhinav/monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
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
}
func (l *Lexer) NextToken() *token.Token {
	skipWhitespace(l)

	var tok token.Token
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: "=="}
		} else {
			tok = token.Token{Type: token.ASSIGN, Literal: "="}
		}
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: "+"}
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: "("}
	case '-':
		tok = token.Token{Type: token.MINUS, Literal: "-"}
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: "!="}
		} else {
			tok = token.Token{Type: token.BANG, Literal: "!"}
		}
	case '/':
		tok = token.Token{Type: token.SLASH, Literal: "/"}
	case '*':
		tok = token.Token{Type: token.ASTERISK, Literal: "*"}
	case '<':
		tok = token.Token{Type: token.LT, Literal: "<"}
	case '>':
		tok = token.Token{Type: token.GT, Literal: ">"}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: ")"}
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: "{"}
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: "}"}
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: ","}
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: ";"}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.ReadKeyword(tok.Literal)
			return &tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return &tok
		} else {
			tok = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
		}
	}
	l.readChar()
	return &tok
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}
func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}
func skipWhitespace(l *Lexer) {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
func isDigit(ch uint8) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}
func isLetter(ch uint8) bool {
	if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_' {
		return true
	}
	return false
}
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	return l.ch
}
