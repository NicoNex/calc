package utils

import (
	"fmt"
	"errors"
	"sync"
)

type Stack struct {
	s []interface{}
	mutex sync.Mutex
}

func NewStack() Stack {
	return Stack{
		make([]interface{}, 0),
		sync.Mutex{},
	}
}

func (s *Stack) Push(n interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.s = append(s.s, n)
}

func (s *Stack) Pop() (interface{}, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var ret interface{}
	var l = len(s.s)

	if l == 0 {
		return nil, errors.New("empty stack")
	}

	ret = s.s[l-1]
	s.s = s.s[:l-1]
	return ret, nil
}

func (s Stack) Peek() (interface{}, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var l = len(s.s)

	if l == 0 {
		return nil, errors.New("empty stack")
	}
	return s.s[l-1], nil
}

func (s Stack) String() string {
	return fmt.Sprintf("stack: %v", s.s)
}
