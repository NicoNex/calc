package ops

type Plus struct {
	l Leaf
	r Leaf
}

func NewPlus(l, r Leaf) Plus {
	return Plus{
		l,
		r,
	}
}

func (p Plus) Eval() float64 {
	return p.l.Eval() + p.r.Eval()
}
