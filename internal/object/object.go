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

type Type uint8

const (
	TypeNull Type = iota
	TypeInteger
	TypeBoolean
	TypeReturnValue
	TypeFunction
	TypeError
)

func (t Type) String() string {
	switch t {
	case TypeInteger:
		return "Integer"
	case TypeBoolean:
		return "Boolean"
	case TypeNull:
		return "Null"
	case TypeReturnValue:
		return "Return Value"
	case TypeFunction:
		return "Function"
	case TypeError:
		return "Error"
	default:
		return "Null"
	}
}

type Object interface {
	Type() Type
	Inspect() string
}
