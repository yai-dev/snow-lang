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

package object

import (
	"bytes"
	"strings"

	"github.com/suenchunyu/snow-lang/internal/ast"
)

type Function struct {
	Env        *Environment
	Body       *ast.BlockStatement
	Parameters []*ast.Identifier
}

func (f *Function) Type() Type {
	return TypeFunction
}

func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := make([]string, 0)
	for _, param := range f.Parameters {
		params = append(params, param.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(") ")
	out.WriteString(f.Body.String())

	return out.String()
}
