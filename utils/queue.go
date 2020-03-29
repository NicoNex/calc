package utils

import "github.com/NicoNex/calc/ops"

type queue []ops.Node

func NewQueue() queue {
	return make(queue)
}

func (q queue) Push(n ops.Node) queue {
	return append(q, n)
}

func (q queue) Pop() (queue, ops.Node) {
	if len(q) == 0 {
		return q, nil
	}
	return q[1:], q[0]
}

func (q queue) Peek() ops.Node {
	if len(q) == 0 {
		return nil
	}
	return p[0]
}
