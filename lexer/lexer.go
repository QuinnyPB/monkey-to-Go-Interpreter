package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string 
	position     int  // curr pos in input (points to currchar)
	readPosition int  // curr reading pos in input (after currchar)
	ch           byte // curr char being examined
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {		
		case '=':
			if l.peekChar() == '=' {
				ch := l.ch
				l.readChar()
				literal := string(ch) + string(l.ch)
				tok = token.Token{Type: token.EQ, Literal: literal}
			} else {
				tok = newToken(token.ASSIGN, l.ch)
			}
		case ';':
			tok = newToken(token.SEMICOLON, l.ch)
		case ':':
			tok = newToken(token.COLON, l.ch)
		case '(':
			tok = newToken(token.LPAREN, l.ch)
		case ')':
			tok = newToken(token.RPAREN, l.ch)
		case ',':
			tok = newToken(token.COMMA, l.ch)
		case '+':
			tok = newToken(token.PLUS, l.ch)
		case '-':
			tok = newToken(token.MINUS, l.ch)
		case '!':
			if l.peekChar() == '=' {
				ch := l.ch
				l.readChar()
				literal := string(ch) + string(l.ch)
				tok = token.Token{Type: token.NOT_EQ, Literal: literal}
			} else {
				tok = newToken(token.BANG, l.ch)
			}
		case '*':
			tok = newToken(token.ASTERISK, l.ch)
		case '/':
			tok = newToken(token.SLASH, l.ch)
		case '<':
			tok = newToken(token.LT, l.ch)
		case '>':
			tok = newToken(token.GT, l.ch)
		case '{':
			tok = newToken(token.LBRACE, l.ch)
		case '}': 	 		
			tok = newToken(token.RBRACE, l.ch)
		case '[':
			tok = newToken(token.LBRACKET, l.ch)
		case ']':
			tok = newToken(token.RBRACKET, l.ch)
		case '"':
			tok.Type = token.STRING
			tok.Literal = l.readString()
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF		
		default:
			if isLetter(l.ch) {
				tok.Literal = l.readIdentifier()
				tok.Type = token.LookupIdent(tok.Literal)
				return tok
			} else if isDigit(l.ch) {
				tok.Type = token.INT
				tok.Literal = l.readNumber()
				return tok
			} else {
				tok = newToken(token.ILLEGAL, l.ch)
			}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input){
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readString() string {
	// transforming string to rune allows for correctly swapping out intended escape characters with their corresponding coutnerparts in go 
	var result []rune

	for {
		l.readChar() // initial read skips over quote

		if l.ch == '"' || l.ch == 0 {
			break
		}

		// if escape char
		if l.ch == '\\' {
			l.readChar()
			switch l.ch {
			case 'a': 		// Alert or bell 
				result = append(result, '\a')
			case 'b':			// Backspace
				result = append(result, '\b')
			case 't':			// Horizontal tab
				result = append(result, '\t')
			case 'n':			// Line feed or newline
				result = append(result, '\n')
			case 'f':			// Form feed
				result = append(result, '\f')
			case 'r':			// Carraige return
				result = append(result, '\r')
			case 'v':   	// Vertical tab
				result = append(result, '\v')
			case '\'':		// Single quote
				result = append(result, '\'')
			case '"':			// double quote
				result = append(result, '"')
			// invalid char. TODO: issue error warning against invalid esc_char
			default:
				result = append(result, '\\', rune(l.ch))
			}
		} else {
			// normal char
			result = append(result, rune(l.ch))		
		}
	}
	
	return string(result)
}
