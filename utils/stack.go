package utils

import (
	"fmt"
	"errors"
)

type Stack struct {
	s []interface{}
}

func NewStack() Stack {
	return Stack{
		make([]interface{}, 0),
	}
}

func (s *Stack) Push(n interface{}) {
	s.s = append(s.s, n)
}

func (s *Stack) Pop() (interface{}, error) {
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
	var l = len(s.s)

	if l == 0 {
		return nil, errors.New("empty stack")
	}
	return s.s[l-1], nil
}

func (s Stack) String() string {
	return fmt.Sprintf("stack: %v", s.s)
}
