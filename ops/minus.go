package ops

type Minus struct {
	l Node
	r Node
}

func NewMinus(l, r Node) Node {
	return Minus{
		l,
		r,
	}
}

func (m Minus) Eval() float64 {
	return m.l.Eval() - m.r.Eval()
}
