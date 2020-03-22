package ops

type Minus struct {
	l Leaf
	r Leaf
}

func NewMinus(l, r Leaf) Minus {
	return Minus{
		l,
		r,
	}
}

func (m Minus) Eval() float64 {
	return m.l.Eval() - m.r.Eval()
}
