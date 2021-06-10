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
	"github.com/suenchunyu/snow-lang/internal/ast"
	"github.com/suenchunyu/snow-lang/internal/token"
)

func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.cur}

	if !p.expectedPeek(token.FlagLP) {
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()

	if !p.expectedPeek(token.FlagLB) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := make([]*ast.Identifier, 0)

	if p.peekTokenIs(token.FlagRP) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := &ast.Identifier{
		Token: p.cur,
		Value: p.cur.Literal,
	}
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(token.FlagComma) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{
			Token: p.cur,
			Value: p.cur.Literal,
		}
		identifiers = append(identifiers, ident)
	}

	if !p.expectedPeek(token.FlagRP) {
		return nil
	}

	return identifiers
}
