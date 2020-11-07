package ast

import "fmt"

type Assign struct {
	l Node
	r Node
}

var vtable map[string]Node

func NewAssign(l, r Node) Node {
	return Assign{
		l,
		r,
	}
}

func (a Assign) Eval() float64 {
	vtable[a.l.String()] = a.r
	return a.r.Eval()
}

func (a Assign) String() string {
	return fmt.Sprintf("(%v=%v)", a.l, a.r)
}

func init() {
	vtable = make(map[string]Node)
}
