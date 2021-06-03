package ast

import "bytes"

type (
    Node interface {
        TokenLiteral() string
        String() string
    }

    Statement interface {
        Node
        statementNode()
    }

    Expression interface {
        Node
        expressionNode()
    }
)

type Program struct {
    Statements []Statement
}

func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    } else {
        return ""
    }
}

func (p *Program) String() string {
    var out bytes.Buffer

    for _, stmt := range p.Statements {
        out.WriteString(stmt.String())
    }

    return out.String()
}