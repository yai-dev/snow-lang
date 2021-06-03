package lexer

import "github.com/suenchunyu/snow-lang/internal/token"

type Lexer struct {
    input string
    pos   int
    rpos  int
    ch    byte
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}

func (l *Lexer) readChar() {
    if l.rpos >= len(l.input) {
        // 'NUL' for ASCII
        l.ch = 0
    } else {
        l.ch = l.input[l.rpos]
    }

    l.pos = l.rpos
    l.rpos += 1
}

func (l *Lexer) peekChar() byte {
    if l.rpos >= len(l.input) {
        return 0
    } else {
        return l.input[l.rpos]
    }
}

func (l *Lexer) NextToken() *token.Token {
    tok := new(token.Token)

    l.skipWhitespace()

    switch l.ch {
    case '=':
        if l.peekChar() == '=' {
            ch := l.ch
            l.readChar()
            tok = &token.Token{
                Flag:    token.FlagEqual,
                Literal: string(ch) + string(l.ch),
            }
        } else {
            tok = token.New(token.FlagAssign, l.ch)
        }
    case '+':
        tok = token.New(token.FlagPlus, l.ch)
    case '-':
        tok = token.New(token.FlagMinus, l.ch)
    case '!':
        if l.peekChar() == '=' {
            ch := l.ch
            l.readChar()
            tok = &token.Token{
                Flag:    token.FlagNotEqual,
                Literal: string(ch) + string(l.ch),
            }
        } else {
            tok = token.New(token.FlagEM, l.ch)
        }

    case '/':
        tok = token.New(token.FlagSlash, l.ch)
    case '*':
        tok = token.New(token.FlagAsterisk, l.ch)
    case ';':
        tok = token.New(token.FlagSemicolon, l.ch)
    case '(':
        tok = token.New(token.FlagLP, l.ch)
    case ')':
        tok = token.New(token.FlagRP, l.ch)
    case ',':
        tok = token.New(token.FlagComma, l.ch)
    case '{':
        tok = token.New(token.FlagLB, l.ch)
    case '}':
        tok = token.New(token.FlagRB, l.ch)
    case '<':
        tok = token.New(token.FlagLT, l.ch)
    case '>':
        tok = token.New(token.FlagGT, l.ch)
    case 0:
        tok.Literal = ""
        tok.Flag = token.FlagEOF
    default:
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier()
            tok.Flag = token.LookupIdent(tok.Literal)
            return tok
        } else if isDigit(l.ch) {
            tok.Literal = l.readNumber()
            tok.Flag = token.FlagInt
            return tok
        } else {
            tok = token.New(token.FlagIllegal, l.ch)
        }
    }

    l.readChar()
    return tok
}

func (l *Lexer) readIdentifier() string {
    pos := l.pos
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[pos:l.pos]
}

func (l *Lexer) readNumber() string {
    pos := l.pos
    for isDigit(l.ch) {
        l.readChar()
    }
    return l.input[pos:l.pos]
}

func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
