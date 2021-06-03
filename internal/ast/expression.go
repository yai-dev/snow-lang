package ast

import "github.com/suenchunyu/snow-lang/internal/token"

type ExpressionStatement struct {
    Token      *token.Token
    Expression Expression
}

func (es *ExpressionStatement) TokenLiteral() string {
    return es.Token.Literal
}

func (es *ExpressionStatement) statementNode() {
    panic("implement me")
}

func (es ExpressionStatement) String() string {
    if es.Expression != nil {
        return es.Expression.String()
    }
    return ""
}