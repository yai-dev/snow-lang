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
                    Flag: token.FlagLet,
                    Literal: "let",
                },
                Name:  &ast.Identifier{
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
