package repl

import (
    "bufio"
    "fmt"
    "io"

    "github.com/suenchunyu/snow-lang/internal/lexer"
    "github.com/suenchunyu/snow-lang/internal/token"
)

const PROMPT = "[Snow Lang]|> "

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)

    for {
        fmt.Printf(PROMPT)
        scanned := scanner.Scan()
        if !scanned {
            return
        }

        line := scanner.Text()
        l := lexer.New(line)

        for tok := l.NextToken(); tok.Flag != token.FlagEOF; tok = l.NextToken() {
            fmt.Printf("%+v\n", tok)
        }
    }
}
