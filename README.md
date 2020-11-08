# calc
A fast and simple cli calculator.

## Install
You can install calc just as any other go project:
```bash
go get -u github.com/NicoNex/calc
```

## Usage
You can use calc in both shell mode and command mode.

### Shell mode
This mode is like the python shell, this mode supports many comfortable and useful terminal features.
```
>>> 2+2
4
>>> 3*(5/(3-4))
-15
>>> my_var = 4*2^5
128
>>> my_var = my_var/2
64
>>> big_var = 2^128
3.402823669209385e+38
>>> 2*/(3+8)
syntax error: invalid operand
	2*/(3+8)
	  ^
```

### Command mode
You can use calc to evaluate expressions with a single command:
```bash
$ calc '3^2+(19-4)*2'
39
$
```

## Supported functions, operators and constants
### Operators
`+`, `-`, `*`, `/`, `^`.
Other operators coming soon.

### Functions
Coming soon.

### Constants
Coming soon.

## Backstory
I was annoyed by the lack of 'bc' on the old Windows version I have to use at work so I thought that writing my own cli calculator would allow me to learn some fun stuff.

## Additional resources
The lexer is heavily inspired by 2011 [Rob Pike's talk on lexical scanning](https://www.youtube.com/watch?v=HxaD_trXwRE&list=LLXhKV860LFdgQVU6kJlRbww&index=2&t=2710s).
The parser implements the Edsger Dijkstra's [Shunting-Yard Algorithm](https://en.wikipedia.org/wiki/Shunting-yard_algorithm) to produce postfix notation from infix.
