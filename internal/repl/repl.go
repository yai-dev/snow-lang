/*
 * Snow-Lang, A Toy-Level Programming Language.
 * Copyright (C) 2021  Suen ChunYu<mailto:sunzhenyucn@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package repl

import (
	"bufio"
	"fmt"
	"io"
	"os/user"

	"github.com/suenchunyu/snow-lang/internal/eval"
	"github.com/suenchunyu/snow-lang/internal/lexer"
	"github.com/suenchunyu/snow-lang/internal/parser"
)

const (
	Prompt = "[Snow Lang]|> "
	Logo   = `  _________                     
 /   _____/ ____   ______  _  __
 \_____  \ /    \ /  _ \ \/ \/ /
 /        \   |  (  <_> )     / 
/_______  /___|  /\____/ \/\_/  
        \/     \/               
`
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	_, _ = io.WriteString(out, Logo)
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	_, _ = io.WriteString(out, fmt.Sprintf("Hello %s! This is the Snow programming language!\n", u.Name))
	_, _ = io.WriteString(out, fmt.Sprintf("Fell free to type in commands!\n"))

	for {
		fmt.Printf(Prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.Parse()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := eval.Eval(program)
		if evaluated != nil {
			_, _ = io.WriteString(out, evaluated.Inspect())
			_, _ = io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	_, _ = io.WriteString(out, Logo)
	_, _ = io.WriteString(out, "Oops! We ran into some awful things here!\n")
	_, _ = io.WriteString(out, " Parser errors: \n")
	for _, errMsg := range errors {
		_, _ = io.WriteString(out, fmt.Sprintf("\t%s\n", errMsg))
	}
}
