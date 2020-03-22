package ops

type Const struct {
	v float64
}

func NewConst(v float64) Const {
	return Const{v}
}

func (c Const) Eval() float64 {
	return c.v
}
