package ops

import "fmt"

type Equal struct {
	name string
	value Node
}

func NewEqual(name string, value Node) Node {
	return Equal{
		name,
		value,
	}
}

func (e Equal) Eval() float64 {
	return e.value.Eval()
}

func (e Equal) String() string {
	return fmt.Sprintf("(%s=%v)", e.name)
}
