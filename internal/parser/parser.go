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

package parser

import (
	"fmt"

	"github.com/suenchunyu/snow-lang/internal/ast"
	"github.com/suenchunyu/snow-lang/internal/lexer"
	"github.com/suenchunyu/snow-lang/internal/token"
)

const (
	_ uint8 = iota
	Lowest
	Equals      // ==
	LessGreater // > or <
	Sum         // +
	Product     // *
	Prefix      // -x or !x
	Call        // customFun(x)
	Index       // array[index]
)

type (
	prefixParseFunc func() ast.Expression
	infixParserFunc func(ast.Expression) ast.Expression
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	cur  *token.Token
	peek *token.Token

	prefix map[token.Flag]prefixParseFunc
	infix  map[token.Flag]infixParserFunc
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: make([]string, 0),
	}

	p.prefix = make(map[token.Flag]prefixParseFunc)
	p.registerPrefix(token.FlagIdent, p.parseIdentifier)
	p.registerPrefix(token.FlagInt, p.parseIntegerLiteral)
	p.registerPrefix(token.FlagEM, p.parsePrefixExpression)
	p.registerPrefix(token.FlagMinus, p.parsePrefixExpression)
	p.registerPrefix(token.FlagTrue, p.parseBoolean)
	p.registerPrefix(token.FlagFalse, p.parseBoolean)
	p.registerPrefix(token.FlagLParen, p.parseGroupedExpression)
	p.registerPrefix(token.FlagIf, p.parseIfExpression)
	p.registerPrefix(token.FlagFunction, p.parseFunctionLiteral)
	p.registerPrefix(token.FlagString, p.parseStringLiteral)
	p.registerPrefix(token.FlagLBracket, p.parseArrayLiteral)

	p.infix = make(map[token.Flag]infixParserFunc)
	p.registerInfix(token.FlagPlus, p.parseInfixExpression)
	p.registerInfix(token.FlagMinus, p.parseInfixExpression)
	p.registerInfix(token.FlagSlash, p.parseInfixExpression)
	p.registerInfix(token.FlagAsterisk, p.parseInfixExpression)
	p.registerInfix(token.FlagEqual, p.parseInfixExpression)
	p.registerInfix(token.FlagNotEqual, p.parseInfixExpression)
	p.registerInfix(token.FlagLessThan, p.parseInfixExpression)
	p.registerInfix(token.FlagGreaterThan, p.parseInfixExpression)
	p.registerInfix(token.FlagLParen, p.parseCallExpression)
	p.registerInfix(token.FlagLBracket, p.parseIndexExpression)

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) registerPrefix(flag token.Flag, fn prefixParseFunc) {
	p.prefix[flag] = fn
}

func (p *Parser) registerInfix(flag token.Flag, fn infixParserFunc) {
	p.infix[flag] = fn
}

var precedences = map[token.Flag]uint8{
	token.FlagEqual:       Equals,
	token.FlagNotEqual:    Equals,
	token.FlagLessThan:    LessGreater,
	token.FlagGreaterThan: LessGreater,
	token.FlagPlus:        Sum,
	token.FlagMinus:       Sum,
	token.FlagSlash:       Product,
	token.FlagAsterisk:    Product,
	token.FlagLParen:      Call,
	token.FlagLBracket:    Index,
}

func (p *Parser) peekPrecedence() uint8 {
	if precedence, ok := precedences[p.peek.Flag]; ok {
		return precedence
	}
	return Lowest
}

func (p *Parser) curPrecedence() uint8 {
	if precedence, ok := precedences[p.cur.Flag]; ok {
		return precedence
	}
	return Lowest
}

func (p *Parser) nextToken() {
	p.cur = p.peek
	p.peek = p.l.NextToken()
}

func (p *Parser) Parse() *ast.Program {
	program := new(ast.Program)
	program.Statements = make([]ast.Statement, 0)

	for p.cur.Flag != token.FlagEOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) curTokenIs(t token.Flag) bool {
	return p.cur.Flag == t
}

func (p *Parser) peekTokenIs(t token.Flag) bool {
	return p.peek.Flag == t
}

func (p *Parser) expectedPeek(t token.Flag) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t token.Flag) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t.String(), p.peek.Flag)
	p.errors = append(p.errors, msg)
}
