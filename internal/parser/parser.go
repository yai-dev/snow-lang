package parser

import (
    "fmt"

    "github.com/suenchunyu/snow-lang/internal/ast"
    "github.com/suenchunyu/snow-lang/internal/lexer"
    "github.com/suenchunyu/snow-lang/internal/token"
)

const (
    _ uint8 = iota
    Lowest
    Equals      // ==
    LessGreater // > or <
    Sum         // +
    Product     // *
    Prefix      // -x or !x
    Call        // customFun(x)
)

type (
    prefixParseFunc func() ast.Expression
    infixParserFunc func(ast.Expression) ast.Expression
)

type Parser struct {
    l      *lexer.Lexer
    errors []string

    cur  *token.Token
    peek *token.Token

    prefix map[token.Flag]prefixParseFunc
    infix  map[token.Flag]infixParserFunc
}

func New(l *lexer.Lexer) *Parser {
    p := &Parser{
        l:      l,
        errors: make([]string, 0),
    }

    p.prefix = make(map[token.Flag]prefixParseFunc)
    p.registerPrefix(token.FlagIdent, p.parseIdentifier)

    p.nextToken()
    p.nextToken()

    return p
}

func (p *Parser) parseIdentifier() ast.Expression {
    return &ast.Identifier{
        Token: p.cur,
        Value: p.cur.Literal,
    }
}

func (p *Parser) registerPrefix(flag token.Flag, fn prefixParseFunc) {
    p.prefix[flag] = fn
}

func (p *Parser) registerInfix(flag token.Flag, fn infixParserFunc) {
    p.infix[flag] = fn
}

func (p *Parser) nextToken() {
    p.cur = p.peek
    p.peek = p.l.NextToken()
}

func (p *Parser) Parse() *ast.Program {
    program := new(ast.Program)
    program.Statements = make([]ast.Statement, 0)

    for p.cur.Flag != token.FlagEOF {
        stmt := p.parseStatement()
        if stmt != nil {
            program.Statements = append(program.Statements, stmt)
        }
        p.nextToken()
    }

    return program
}

func (p *Parser) Errors() []string {
    return p.errors
}

func (p *Parser) parseStatement() ast.Statement {
    switch p.cur.Flag {
    case token.FlagLet:
        return p.parseLetStatement()
    case token.FlagReturn:
        return p.parseReturnStatement()
    default:
        return p.parseExpressionStatement()
    }
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
    stmt := &ast.LetStatement{Token: p.cur}

    if !p.expectedPeek(token.FlagIdent) {
        return nil
    }

    stmt.Name = &ast.Identifier{
        Token: p.cur,
        Value: p.cur.Literal,
    }

    if !p.expectedPeek(token.FlagAssign) {
        return nil
    }

    for !p.curTokenIs(token.FlagSemicolon) {
        p.nextToken()
    }

    return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
    stmt := &ast.ReturnStatement{Token: p.cur}

    p.nextToken()

    for !p.curTokenIs(token.FlagSemicolon) {
        p.nextToken()
    }
    return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
    stmt := &ast.ExpressionStatement{Token: p.cur}

    stmt.Expression = p.parseExpression(Lowest)

    if p.peekTokenIs(token.FlagSemicolon) {
        p.nextToken()
    }

    return stmt
}

func (p *Parser) parseExpression(precedence uint8) ast.Expression {
    prefix := p.prefix[p.cur.Flag]
    if prefix == nil {
        return nil
    }

    left := prefix()

    return left
}

func (p *Parser) curTokenIs(t token.Flag) bool {
    return p.cur.Flag == t
}

func (p *Parser) peekTokenIs(t token.Flag) bool {
    return p.peek.Flag == t
}

func (p *Parser) expectedPeek(t token.Flag) bool {
    if p.peekTokenIs(t) {
        p.nextToken()
        return true
    } else {
        p.peekError(t)
        return false
    }
}

func (p *Parser) peekError(t token.Flag) {
    msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peek.Flag)
    p.errors = append(p.errors, msg)
}
