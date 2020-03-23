package ops

type Variable struct {
	v Node
}

func NewVariable(v Node) Node {
	return Variable{
		v,
	}
}

func (v Variable) Eval() float64 {
	return v.v.Eval()
}
