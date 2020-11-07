package ast

type Variable struct {
	n string
}

func NewVariable(n string) Node {
	return Variable{
		n,
	}
}

func (v Variable) Eval() float64 {
	if n, ok := vtable[v.n]; ok {
		return n.Eval()
	}
	return 0
}

func (v Variable) String() string {
	return v.n
}
