# Snow Lang
[![Go Report Card](https://goreportcard.com/badge/github.com/suenchunyu/snow-lang)](https://goreportcard.com/report/github.com/suenchunyu/snow-lang)


```text
                                      _________                     
                                    /   _____/ ____   ______  _  __
                                    \_____  \ /    \ /  _ \ \/ \/ /
                                    /        \   |  (  <_> )     /
                                    /_______  /___|  /\____/ \/\_/  
                                    \/     \/
```

`Snow Lang` is a **Toy-Level** programming language writen in Go and implemented an interpreter for it.

## Example for Snow Lang

```text
let version = 1;
let name = "Snow programming language";
let arr = [1, 2, 3];
let bool_val = true;

let awesome_val = 1997 * 1110;
let awesome_arr_val = [1 + 1, 2 + 2, 3 * 3];

let fibonacci = fn(x) {
  if (x == 0) {
    0
  } else {
    if (x == 1) {
      return 1;
    } else {
      fibonacci(x - 1) + fibonacci(x - 2);
    }
  }
};

let result = fibonacci(1);

let persons = [{"name": "SuenChunYu", "age": 24}];
```

## Why named 'Snow Lang'?

It' simple and crystal, `Snow` is the homonym of snowflakes in Chinese(`雪花`), and the `雪花` is homophonic for my
girlfriend's name in Chinese.

## Can I use `Snow Lang` for production?

Of course not. `Snow Lang` is just toy and I learn the compilers principle through it.

## Roadmap

- [x] Lexing and defining tokens.
- [x] Simple REPL.
- [x] Parsing `let` statements.
- [x] Parsing `return` statements.
- [x] Parsing expressions.
- [ ] Pratt Parsing algorithm.
- [ ] Read parse print loop.
- [ ] Evaluation.
- [ ] Interpreter Extending.
- [ ] Makefile build script.
- [ ] Evaluation codes from `*.snow` files.
- [ ] Releasing CI/CD scripts.

> Yeah, Long way to go. :)

## Reference

- [Writing An Interpreter In Go](https://interpreterbook.com/)
