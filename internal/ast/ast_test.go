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

package ast_test

import (
	"testing"

	"github.com/suenchunyu/snow-lang/internal/ast"
	"github.com/suenchunyu/snow-lang/internal/token"
)

func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: &token.Token{
					Flag:    token.FlagLet,
					Literal: "let",
				},
				Name: &ast.Identifier{
					Token: &token.Token{
						Flag:    token.FlagIdent,
						Literal: "foo",
					},
					Value: "foo",
				},
				Value: &ast.Identifier{
					Token: &token.Token{
						Flag:    token.FlagIdent,
						Literal: "bar",
					},
					Value: "bar",
				},
			},
		},
	}

	if program.String() != "let foo = bar;" {
		t.Errorf("program.String wrong. got = %q", program.String())
	}
}
