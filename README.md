# calc
Cli calculator for educational purpose.

The lexer is heavily inspired by 2011 [Rob Pike's talk on lexical scanning](https://www.youtube.com/watch?v=HxaD_trXwRE&list=LLXhKV860LFdgQVU6kJlRbww&index=2&t=2710s).

The parser implements the Edsger Dijkstra's [Shunting-Yard Algorithm](https://en.wikipedia.org/wiki/Shunting-yard_algorithm) to parse infix notation and produce postfix notation.
