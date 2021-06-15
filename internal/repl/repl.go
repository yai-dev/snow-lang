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
	"runtime"

	"github.com/suenchunyu/snow-lang/internal/eval"
	"github.com/suenchunyu/snow-lang/internal/lexer"
	"github.com/suenchunyu/snow-lang/internal/object"
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

var (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
	white  = "\033[97m"
)

func init() {
	if runtime.GOOS == "windows" {
		reset = ""
		red = ""
		green = ""
		yellow = ""
		blue = ""
		purple = ""
		cyan = ""
		gray = ""
		white = ""
	}
}

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnv()

	_, _ = io.WriteString(out, green+Logo+reset)
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	_, _ = io.WriteString(out, blue+fmt.Sprintf("Hello %s! This is the Snow programming language!\n", u.Name)+reset)
	_, _ = io.WriteString(out, blue+fmt.Sprintf("Fell free to type in commands!\n")+reset)

	for {
		fmt.Printf(yellow + Prompt + reset)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.Parse()
		if len(p.Errors()) != 0 {
			printErrors(out, p.Errors())
			continue
		}

		evaluated := eval.Eval(program, env)
		if evaluated != nil {
			if evaluated.Type() == object.TypeError {
				printErrors(out, []string{evaluated.Inspect()})
				continue
			}
			_, _ = io.WriteString(out, evaluated.Inspect())
			_, _ = io.WriteString(out, "\n")
		}
	}
}

func printErrors(out io.Writer, errors []string) {
	_, _ = io.WriteString(out, red+"Oops! We ran into some awful things here!\n"+reset)
	for _, errMsg := range errors {
		_, _ = io.WriteString(out, red+fmt.Sprintf("\t%s\n", errMsg)+reset)
	}
}
