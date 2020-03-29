package utils

import "github.com/NicoNex/calc/ops"

type stack []ops.Node

func NewQueue() stack {
	return make(stack)
}

func (s stack) Push(n ops.Node) stack {
	return append(s, n)
}

func (s stack) Pop() (stack, ops.Node) {
	var l = len(s)

	if l == 0 {
		return s, nil
	}
	return s[:l-1], s[l-1]
}

func (s stack) Peek() ops.Node {
	var l = len(s)

	if l == 0 {
		return nil
	}
	return p[l-1]
}
