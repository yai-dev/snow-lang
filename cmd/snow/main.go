package main

import (
    "fmt"
    "os"
    "os/user"

    "github.com/suenchunyu/snow-lang/internal/repl"
)

func main() {
    u, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Hello %s! This is the Snow programming language!\n", u.Name)
    fmt.Printf("Fell free to type in commands!\n")
    repl.Start(os.Stdin, os.Stdout)
}
