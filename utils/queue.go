package utils

import (
	"fmt"
	"errors"
	"sync"
)

type Queue struct {
	q []interface{}
	mutex sync.Mutex
}

func NewQueue() Queue {
	return Queue{
		make([]interface{}, 0),
		sync.Mutex{},
	}
}

func (q *Queue) Push(n interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.q = append(q.q, n)
}

func (q *Queue) Pop() (interface{}, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	var ret interface{}

	if len(q.q) == 0 {
		return nil, NewEmptyQueue()
	}

	ret = q.q[0]
	q.q = q.q[1:]
	return ret, nil
}

func (q Queue) Peek() (interface{}, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if len(q.q) == 0 {
		return nil, NewEmptyQueue()
	}
	return q.q[0], nil
}

func (q Queue) String() string {
	return fmt.Sprintf("queue: %v", q.q)
}
