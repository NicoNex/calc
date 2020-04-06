package parser

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type itemType int

const (
	operator itemType = iota
	operand
	variable
	bracket
	itemError
)

type item struct {
	typ itemType
	val string
}

type lexer struct {
	input string
	start int
	pos   int
	width int
	prev  itemType
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
	return fmt.Sprintf("%q", i.val)
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
	l.pos -= l.width
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

// Consumes the next rune if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}

	l.backup()
	return false
}

// Consumes all the runes if they're in the valid set.
func (l *lexer) acceptRun(valid string) bool {
	for strings.IndexRune(valid, l.next()) >= 0 {

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
	l.prev = t
}

// This is the main loop that makes the state machine work.
func (l *lexer) run() {
	for state := lexExpression; state != nil; {
		state = state(l)
	}
	close(l.items)
}

func (l *lexer) printState() {
	fmt.Printf(
		"st: %d\npos: %d\nwh: %d\ncur: %s\nin: %s\n\n",
		l.start, l.pos, l.width, string(l.peek()), l.input,
	)
}

func lexOperator(l *lexer) stateFn {
	l.accept("+-*/=")
	l.emit(operator)
	return lexExpression
}

func lexBracket(l *lexer) stateFn {
	l.accept("()")
	l.emit(bracket)
	return lexExpression
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

	l.emit(operand)
	return lexExpression
}

func lexVariable(l *lexer) stateFn {
	var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if l.acceptRun(chars) {
		l.emit(variable)
	}

	return lexExpression
}

func isLetter(r rune) bool {
	return r == '_' || unicode.IsLetter(r)
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isOperator(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/' || r == '='
}

func isBracket(r rune) bool {
	return r == '(' || r == ')'
}

func lexExpression(l *lexer) stateFn {
	for {
		switch r := l.next(); {

		case isSpace(r):
			l.ignore()

		case r == '\n':
			return nil

		case isLetter(r):
			l.backup()
			return lexVariable

		case isOperator(r):
			l.backup()
			switch l.prev {
			case operator:
				return lexNumber
			default:
				return lexOperator
			}

		case isBracket(r):
			l.backup()
			return lexBracket

		default:
			l.backup()
			return lexNumber
		}
	}

	return nil
}

// Returns the lexer object and the channel used to receive the lexed items.
func lex(input string) (*lexer, chan item) {
	var l = &lexer{
		input: input,
		items: make(chan item),
	}

	go l.run()
	return l, l.items
}
