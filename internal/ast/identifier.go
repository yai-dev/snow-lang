package ast

import "github.com/suenchunyu/snow-lang/internal/token"

type Identifier struct {
    Token *token.Token
    Value string
}

func (i *Identifier) TokenLiteral() string {
    return i.Token.Literal
}

func (i *Identifier) expressionNode() {
    panic("implement me")
}

func (i Identifier) String() string {
    return i.Value
}

