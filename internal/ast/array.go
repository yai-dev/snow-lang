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

package ast

import (
	"bytes"
	"strings"

	"github.com/suenchunyu/snow-lang/internal/token"
)

type ArrayLiteral struct {
	Token    *token.Token
	Elements []Expression
}

func (al *ArrayLiteral) TokenLiteral() string {
	return al.Token.Literal
}

func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	elements := make([]string, 0)
	for _, ele := range al.Elements {
		elements = append(elements, ele.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

func (al *ArrayLiteral) expressionNode() {
	panic("implement me")
}
