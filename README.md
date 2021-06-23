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

## Syntax Examples

> The semicolon(`;`) is optional like `JavaScript`.

### 1. Declare and Assignment

```javascript
let a = 1;
a = 2;
let str = "Hello, Snow";
let bool = true;
```

### 2. Loop with `for`

```javascript
let sum = 0;

for (let i = 0; i < 50; i = i + 1) {
    sum = sum + i;
}
```

### 3. Condition

```javascript
let a = true;
let b = 0;

if (a) {
    b = 1;
} else {
    b = 2;
}
```

### 4. Function

```javascript
fn add(x, y) {
    // `return` is optional
  x + y
}

add(1, 2);

fn fibonacci(num) {
    // `bracket` is optional 
    if num < 2 {
        return num
    } else {
        return fibonacci(num - 1) + fibonacci(num - 2)
    }
}

fibonacci(5);
```

### 5. Array and Maps

```javascript
let arr = [1, "2", 3 + 4, 5];
let num = arr[2];
arr[3] = 6;
arr[2] = 1 + 7;

let map = {"name": "Jack", "Age": 24, "Gender": "Male"};
let name = map["name"];
let age = map["age"];
map["age"] = 23;
map["newest_field"] = "I am newest field";
```

### 6. Mathematical operations

```javascript
let a = 1 + 1;
let b = 1 - 1;
let c = 1 * 1;
let d = 1 / 1;
let e = a + b;
let f = b - c;
let g = true;
let h = false;
let i = true == false
let j = 1 < 2;
let k = 1 >= 2;
let l = g == h;
let m = g != h;
```

### 7. Built-in Functions

```javascript
let str = "Hello, Snow!";
let arr = [1, 1, 1, 0];
let map = {"Name": "Snow", "Age": 27, "Gender": "Femal"};

len(str); // 13
len(arr); // 4
head(str); // "H"
head(arr); // 1
tail(str); // "!"
tail(arr); // 0
rest(str); // "ello, Snow"
rest(arr); // [1, 1]
push(str, "!"); // "ello, Snow!"
push(arr, 1); // [1, 1, 1]
print(str, "world"); // "ello, Snow!"\n"world"
print(arr) // [1, 1, 1]
timestamp() // milliseconds since '1970-01-01 00:00:00 UTC'
```

## Why named 'Snow Lang'?

It' simple and crystal, `Snow` is the homonym of snowflakes in Chinese(`雪花`), and the `雪花` is homophonic for my
girlfriend's name in Chinese.

## Can I use `Snow Lang` for production?

Of course not. `Snow Lang` is just toy and I learn the compilers principle through it.

## Roadmap

**Finished:**

- [x] Lexing and defining tokens.
- [x] Simple REPL.
- [x] Parsing `let` statements.
- [x] Parsing `return` statements.
- [x] Parsing expressions.
- [x] Parsing grouped expressions.
- [x] Parsing Boolean
- [x] Parsing block statements.
- [x] Parsing `if` / `if else` statements.
- [x] Parsing `fn` literal.
- [x] Parsing function calling expression.
- [x] Read parse print loop.
- [x] Evaluation.
  - [x] Integer evaluation.
  - [x] Boolean evaluation.
  - [x] Bang(`!`) operator evaluation.
  - [x] `if` statements evaluation.
  - [x] Error handling.
  - [x] Environment & Bindings.
  - [x] Functions & Functions calling.

**WIP:**

- [ ] Interpreter Extending.
  - [x] Strings
  - [x] Built-in Function: `len()`
  - [x] Array
    - [x] Parsing array literal
    - [x] Support index operation
    - [ ] Evaluating array literals
  - [ ] Maps
  - [ ] Built-in Function: `head()`
  - [ ] Built-in Function: `tail()`
  - [ ] Built-in Function: `rest()`
  - [ ] Built-in Function: `push()`
  - [ ] Built-in Function: `print()`
  - [ ] Built-in Function: `timestamp()`
- [ ] Makefile build script.
- [ ] Evaluation codes from `*.snow` files.
- [ ] Releasing CI/CD scripts.

> Yeah, Long way to go. :)

## Reference

- [Writing An Interpreter In Go](https://interpreterbook.com/)
