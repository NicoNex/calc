package parser

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type itemType int
const (
	operator itemType = iota
	constant
	variable
	itemError
)

type item struct {
	typ itemType
	val string
}

type lexer struct {
	input string
	start int
	pos int
	width int
	items chan item
}

type stateFn func(*lexer) stateFn

func (i item) String() string {
	switch i.typ {
	case itemError:
		return i.val
	}
	if len(i.val) > 10 {
		return fmt.Sprintf("%.10q...", i.val)
	}
	return fmt.Sprintf("%q", i)
}

func (l *lexer) next() (rune rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return '\n'
	}

	rune, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return rune
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) backup() {
	l.start -= l.width
}

func (l *lexer) peek() rune {
	rune := l.next()
	l.backup()
	return rune
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{
		itemError,
		fmt.Sprintf(format, args...),
	}
	return nil
}

// accepts consumes the next rune
// if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}

	l.backup()
	return false
}

func (l *lexer) acceptRun(valid string) bool {
	for strings.IndexRune(valid, l.next()) >= 0 {
		;
	}
	l.backup()
	return true
}

func (l *lexer) emit(t itemType) {
	l.items <- item{
		t,
		l.input[l.start:l.pos],
	}
	l.start = l.pos
}

func (l *lexer) run() {
	for state := lexActionState; state != nil; {
		state = state(l)
	}
	close(l.items)
}

func lexOperator(l *lexer) stateFn {
	var digits = "+-*/="
	l.accept(digits)
	l.emit(operator)
	return lexNumber
}

func lexNumber(l *lexer) stateFn {
	var digits = "0123456789"

	// Optional leading sign
	l.accept("+-")

	// Is it hex?
	if l.accept("0") && l.accept("xX") {
		digits = "0123456789abcdefABCDEF"
	}

	l.acceptRun(digits)
	if l.accept(".") {
		l.acceptRun(digits)
	}

	if l.accept("eE") {
		l.accept("+-")
		l.accept("0123456789")
	}

	l.emit(constant)
	return lexOperator
}

func lexVariable(l *lexer) stateFn {
	return nil
}

func isAlphanumeric(r rune) bool {
	return false
}

func lexActionState(l *lexer) stateFn {
	for {
		var r = l.peek()

		if r == '\n' {
			return nil
		} else if isAlphanumeric(r) {
			return lexVariable
		} else {
			return lexNumber
		}
	}

	return nil
}

func lex(input string) (*lexer, chan item) {
	var l = &lexer{
		input: input,
		items: make(chan item),
	}

	go l.run()
	return l, l.items
}
