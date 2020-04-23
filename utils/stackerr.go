package utils

type EmptyStack struct {
	msg string
}

func NewEmptyStack() error {
	return EmptyStack{"error: empty stack"}
}

func (s EmptyStack) Error() string {
	return s.msg
}
