package ops

type Variable struct {
	v Leaf
}

func NewVariable(v Leaf) Variable {
	return Variable{
		v,
	}
}

func (v Variable) Eval() float64 {
	return v.v.Eval()
}
