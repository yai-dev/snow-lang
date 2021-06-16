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

package parser_test

import (
	"testing"

	"github.com/suenchunyu/snow-lang/internal/ast"
	"github.com/suenchunyu/snow-lang/internal/lexer"
	"github.com/suenchunyu/snow-lang/internal/parser"
)

func TestStringLiteralExpression(t *testing.T) {
	input := `"hello world";`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.Parse()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	literal, ok := stmt.Expression.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("exp not *ast.StringLiteral. got = %T", stmt.Expression)
	}

	if literal.Value != "hello world" {
		t.Errorf("literal.Value not %q. got = %q", "hello world", literal.Value)
	}
}
