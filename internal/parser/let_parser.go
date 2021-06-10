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

	p.nextToken()

	stmt.Value = p.parseExpression(Lowest)

	for p.curTokenIs(token.FlagSemicolon) {
		p.nextToken()
	}

	return stmt
}
