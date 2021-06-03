package ast

import (
    "bytes"

    "github.com/suenchunyu/snow-lang/internal/token"
)

type LetStatement struct {
    Token *token.Token
    Name  *Identifier
    Value Expression
}

func (ls *LetStatement) TokenLiteral() string {
    return ls.Token.Literal
}

func (ls *LetStatement) statementNode() {
    panic("implement me")
}

func (ls LetStatement) String() string {
    var out bytes.Buffer

    out.WriteString(ls.TokenLiteral() + " ")
    out.WriteString(ls.Name.String())
    out.WriteString(" = ")

    if ls.Value != nil {
        out.WriteString(ls.Value.String())
    }

    out.WriteString(";")

    return out.String()
}