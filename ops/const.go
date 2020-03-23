package ops

type Const struct {
	v float64
}

func NewConst(v float64) Node {
	return Const{v}
}

func (c Const) Eval() float64 {
	return c.v
}
