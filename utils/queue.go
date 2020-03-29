package utils

import (
	"fmt"
	"errors"
)

type Queue struct {
	q []interface{}
}

func NewQueue() Queue {
	return Queue{
		make([]interface{}, 0),
	}
}

func (q *Queue) Push(n interface{}) {
	 q.q = append(q.q, n)
}

func (q *Queue) Pop() (interface{}, error) {
	var ret interface{}

	if len(q.q) == 0 {
		return nil, errors.New("empty queue")
	}

	ret = q.q[0]
	q.q = q.q[1:]
	return ret, nil
}

func (q Queue) Peek() (interface{}, error) {
	if len(q.q) == 0 {
		return nil, errors.New("empty queue")
	}
	return q.q[0], nil
}

func (q Queue) String() string {
	return fmt.Sprintf("queue: %v", q.q)
}
