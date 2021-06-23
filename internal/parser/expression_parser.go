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
	"github.com/suenchunyu/snow-lang/internal/token"
)

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.cur}

	stmt.Expression = p.parseExpression(Lowest)

	if p.peekTokenIs(token.FlagSemicolon) {
		p.nextToken()
	}

	return stmt
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

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(Lowest)

	if !p.expectedPeek(token.FlagRParen) {
		return nil
	}
	return exp
}

func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.cur}

	if !p.expectedPeek(token.FlagLParen) {
		return nil
	}

	p.nextToken()
	expression.Condition = p.parseExpression(Lowest)

	if !p.expectedPeek(token.FlagRParen) {
		return nil
	}

	if !p.expectedPeek(token.FlagLBrace) {
		return nil
	}

	expression.Consequence = p.parseBlockStatement()

	if p.peekTokenIs(token.FlagElse) {
		p.nextToken()

		if !p.expectedPeek(token.FlagLBrace) {
			return nil
		}

		expression.Alternative = p.parseBlockStatement()
	}
	return expression
}

func (p *Parser) parseExpression(precedence uint8) ast.Expression {
	prefix := p.prefix[p.cur.Flag]
	if prefix == nil {
		msg := fmt.Sprintf("no prefix parse function for %s found", p.cur.Flag.String())
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

func (p *Parser) parseExpressionList(end token.Flag) []ast.Expression {
	list := make([]ast.Expression, 0)

	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()
	list = append(list, p.parseExpression(Lowest))

	for p.peekTokenIs(token.FlagComma) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(Lowest))
	}

	if !p.expectedPeek(end) {
		return nil
	}
	return list
}
