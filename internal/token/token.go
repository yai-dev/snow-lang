package token

type Flag int8

const (
    FlagIllegal Flag = iota
    FlagEOF

    FlagIdent
    FlagInt

    FlagAssign
    FlagPlus
    FlagMinus
    FlagEM
    FlagAsterisk
    FlagSlash

    FlagComma
    FlagSemicolon

    FlagLP
    FlagRP
    FlagLB
    FlagRB
    FlagLT
    FlagGT

    FlagEqual
    FlagNotEqual

    FlagFunction
    FlagLet
    FlagTrue
    FlagFalse
    FlagIf
    FlagElse
    FlagReturn
)

func (f Flag) String() string {
    switch f {
    case FlagIllegal:
        return "ILLEGAL"
    case FlagEOF:
        return "EOF"
    case FlagIdent:
        return "IDENT"
    case FlagInt:
        return "INT"
    case FlagAssign:
        return "="
    case FlagPlus:
        return "+"
    case FlagMinus:
        return "-"
    case FlagEM:
        return "!"
    case FlagAsterisk:
        return "*"
    case FlagSlash:
        return "/"
    case FlagComma:
        return ","
    case FlagSemicolon:
        return ";"
    case FlagLP:
        return "("
    case FlagRP:
        return ")"
    case FlagLB:
        return "{"
    case FlagRB:
        return "}"
    case FlagLT:
        return "<"
    case FlagGT:
        return ">"
    case FlagEqual:
        return "=="
    case FlagNotEqual:
        return "!="
    case FlagFunction:
        return "FUNCTION"
    case FlagLet:
        return "LET"
    case FlagTrue:
        return "TRUE"
    case FlagFalse:
        return "FALSE"
    case FlagIf:
        return "IF"
    case FlagElse:
        return "ELSE"
    case FlagReturn:
        return "RETURN"
    default:
        return "ILLEGAL"
    }
}

var keywords = map[string]Flag{
    "fn":  FlagFunction,
    "let": FlagLet,
    "true": FlagTrue,
    "false": FlagFalse,
    "if": FlagIf,
    "else": FlagElse,
    "return": FlagReturn,
}

type Token struct {
    Flag    Flag
    Literal string
}

func New(flag Flag, ch byte) *Token {
    return &Token{
        Flag:    flag,
        Literal: string(ch),
    }
}

func LookupIdent(ident string) Flag {
    if flag, ok := keywords[ident]; ok {
        return flag
    }
    return FlagIdent
}