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

package lexer_test

import (
	"testing"

	"github.com/suenchunyu/snow-lang/internal/lexer"
	"github.com/suenchunyu/snow-lang/internal/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
  return true;
} else {
  return false;
}
5 != 10;
5 == 10;
"foobar"
"foo bar"
`

	tests := []struct {
		expectedFlag    token.Flag
		expectedLiteral string
	}{
		{token.FlagLet, "let"},
		{token.FlagIdent, "five"},
		{token.FlagAssign, "="},
		{token.FlagInt, "5"},
		{token.FlagSemicolon, ";"},
		{token.FlagLet, "let"},
		{token.FlagIdent, "ten"},
		{token.FlagAssign, "="},
		{token.FlagInt, "10"},
		{token.FlagSemicolon, ";"},
		{token.FlagLet, "let"},
		{token.FlagIdent, "add"},
		{token.FlagAssign, "="},
		{token.FlagFunction, "fn"},
		{token.FlagLP, "("},
		{token.FlagIdent, "x"},
		{token.FlagComma, ","},
		{token.FlagIdent, "y"},
		{token.FlagRP, ")"},
		{token.FlagLB, "{"},
		{token.FlagIdent, "x"},
		{token.FlagPlus, "+"},
		{token.FlagIdent, "y"},
		{token.FlagSemicolon, ";"},
		{token.FlagRB, "}"},
		{token.FlagSemicolon, ";"},
		{token.FlagLet, "let"},
		{token.FlagIdent, "result"},
		{token.FlagAssign, "="},
		{token.FlagIdent, "add"},
		{token.FlagLP, "("},
		{token.FlagIdent, "five"},
		{token.FlagComma, ","},
		{token.FlagIdent, "ten"},
		{token.FlagRP, ")"},
		{token.FlagSemicolon, ";"},
		{token.FlagEM, "!"},
		{token.FlagMinus, "-"},
		{token.FlagSlash, "/"},
		{token.FlagAsterisk, "*"},
		{token.FlagInt, "5"},
		{token.FlagSemicolon, ";"},
		{token.FlagInt, "5"},
		{token.FlagLT, "<"},
		{token.FlagInt, "10"},
		{token.FlagGT, ">"},
		{token.FlagInt, "5"},
		{token.FlagSemicolon, ";"},
		{token.FlagIf, "if"},
		{token.FlagLP, "("},
		{token.FlagInt, "5"},
		{token.FlagLT, "<"},
		{token.FlagInt, "10"},
		{token.FlagRP, ")"},
		{token.FlagLB, "{"},
		{token.FlagReturn, "return"},
		{token.FlagTrue, "true"},
		{token.FlagSemicolon, ";"},
		{token.FlagRB, "}"},
		{token.FlagElse, "else"},
		{token.FlagLB, "{"},
		{token.FlagReturn, "return"},
		{token.FlagFalse, "false"},
		{token.FlagSemicolon, ";"},
		{token.FlagRB, "}"},
		{token.FlagInt, "5"},
		{token.FlagNotEqual, "!="},
		{token.FlagInt, "10"},
		{token.FlagSemicolon, ";"},
		{token.FlagInt, "5"},
		{token.FlagEqual, "=="},
		{token.FlagInt, "10"},
		{token.FlagSemicolon, ";"},
		{token.FlagString, "foobar"},
		{token.FlagString, "foo bar"},
		{token.FlagEOF, ""},
	}

	l := lexer.New(input)

	for idx, tt := range tests {
		tok := l.NextToken()

		if tok.Flag != tt.expectedFlag {
			t.Fatalf("tests[%d] - wrong token type, expected = %q, got = %q", idx, tt.expectedFlag, tok.Flag)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - wrong literal, expected = %q, got = %q", idx, tt.expectedLiteral, tok.Literal)
		}
	}
}
