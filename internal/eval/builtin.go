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

package eval

import "github.com/suenchunyu/snow-lang/internal/object"

var builtin = map[string]*object.Builtin{
	"len": {builtinFunctionLen()},
}

func builtinFunctionLen() object.BuiltinFunction {
	return func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return throw("wrong number of arguments. got %d, want 1", len(args))
		}

		switch arg := args[0].(type) {
		case *object.String:
			return &object.Integer{Value: int64(len(arg.Value))}
		default:
			return throw("argument type to `len` not supported")
		}
	}
}
