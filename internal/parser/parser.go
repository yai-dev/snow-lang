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
	"strconv"

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

	p.infix = make(map[token.Flag]infixParserFunc)
	p.registerInfix(token.FlagPlus, p.parseInfixExpression)
	p.registerInfix(token.FlagMinus, p.parseInfixExpression)
	p.registerInfix(token.FlagSlash, p.parseInfixExpression)
	p.registerInfix(token.FlagAsterisk, p.parseInfixExpression)
	p.registerInfix(token.FlagEqual, p.parseInfixExpression)
	p.registerInfix(token.FlagNotEqual, p.parseInfixExpression)
	p.registerInfix(token.FlagLT, p.parseInfixExpression)
	p.registerInfix(token.FlagGT, p.parseInfixExpression)

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.cur,
		Value: p.cur.Literal,
	}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.cur}

	value, err := strconv.ParseInt(p.cur.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.cur.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.cur,
		Operator: p.cur.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(Prefix)

	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.cur,
		Left:     left,
		Operator: p.cur.Literal,
	}

	precedences := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedences)

	return expression
}

func (p *Parser) registerPrefix(flag token.Flag, fn prefixParseFunc) {
	p.prefix[flag] = fn
}

func (p *Parser) registerInfix(flag token.Flag, fn infixParserFunc) {
	p.infix[flag] = fn
}

var precedences = map[token.Flag]uint8{
	token.FlagEqual:    Equals,
	token.FlagNotEqual: Equals,
	token.FlagLT:       LessGreater,
	token.FlagGT:       LessGreater,
	token.FlagPlus:     Sum,
	token.FlagMinus:    Sum,
	token.FlagSlash:    Product,
	token.FlagAsterisk: Product,
	token.FlagLP:       Call,
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

func (p *Parser) parseStatement() ast.Statement {
	switch p.cur.Flag {
	case token.FlagLet:
		return p.parseLetStatement()
	case token.FlagReturn:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.cur}

	if !p.expectedPeek(token.FlagIdent) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.cur,
		Value: p.cur.Literal,
	}

	if !p.expectedPeek(token.FlagAssign) {
		return nil
	}

	for !p.curTokenIs(token.FlagSemicolon) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.cur}

	p.nextToken()

	for !p.curTokenIs(token.FlagSemicolon) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.cur}

	stmt.Expression = p.parseExpression(Lowest)

	if p.peekTokenIs(token.FlagSemicolon) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence uint8) ast.Expression {
	prefix := p.prefix[p.cur.Flag]
	if prefix == nil {
		msg := fmt.Sprintf("no prefix parse function for %s found", p.cur.Flag)
		p.errors = append(p.errors, msg)
		return nil
	}

	left := prefix()

	for !p.peekTokenIs(token.FlagSemicolon) && precedence < p.peekPrecedence() {
		infix := p.infix[p.peek.Flag]
		if infix == nil {
			return left
		}

		p.nextToken()

		left = infix(left)
	}

	return left
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
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peek.Flag)
	p.errors = append(p.errors, msg)
}
