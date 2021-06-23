/*
 * Snow-Lang, A Toy-Level Programming Language.
 * Copyright (C) 2021  Suen ChunYu<mailto:sunzhenyucn@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package lexer

import "github.com/suenchunyu/snow-lang/internal/token"

type Lexer struct {
	input string
	pos   int
	rpos  int
	ch    byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.rpos >= len(l.input) {
		// 'NUL' for ASCII
		l.ch = 0
	} else {
		l.ch = l.input[l.rpos]
	}

	l.pos = l.rpos
	l.rpos += 1
}

func (l *Lexer) peekChar() byte {
	if l.rpos >= len(l.input) {
		return 0
	} else {
		return l.input[l.rpos]
	}
}

func (l *Lexer) NextToken() *token.Token {
	tok := new(token.Token)

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = &token.Token{
				Flag:    token.FlagEqual,
				Literal: string(ch) + string(l.ch),
			}
		} else {
			tok = token.New(token.FlagAssign, l.ch)
		}
	case '+':
		tok = token.New(token.FlagPlus, l.ch)
	case '-':
		tok = token.New(token.FlagMinus, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = &token.Token{
				Flag:    token.FlagNotEqual,
				Literal: string(ch) + string(l.ch),
			}
		} else {
			tok = token.New(token.FlagEM, l.ch)
		}

	case '/':
		tok = token.New(token.FlagSlash, l.ch)
	case '*':
		tok = token.New(token.FlagAsterisk, l.ch)
	case ';':
		tok = token.New(token.FlagSemicolon, l.ch)
	case '(':
		tok = token.New(token.FlagLParen, l.ch)
	case ')':
		tok = token.New(token.FlagRParen, l.ch)
	case ',':
		tok = token.New(token.FlagComma, l.ch)
	case '{':
		tok = token.New(token.FlagLBrace, l.ch)
	case '}':
		tok = token.New(token.FlagRBrace, l.ch)
	case '<':
		tok = token.New(token.FlagLessThan, l.ch)
	case '>':
		tok = token.New(token.FlagGreaterThan, l.ch)
	case '[':
		tok = token.New(token.FlagLBracket, l.ch)
	case ']':
		tok = token.New(token.FlagRBracket, l.ch)
	case '"':
		tok.Flag = token.FlagString
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Flag = token.FlagEOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Flag = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Flag = token.FlagInt
			return tok
		} else {
			tok = token.New(token.FlagIllegal, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	pos := l.pos
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readString() string {
	position := l.pos + 1
	for {
		l.readChar()
		if l.ch == '"' {
			break
		}
	}
	return l.input[position:l.pos]
}

func (l *Lexer) readNumber() string {
	pos := l.pos
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
