package ops

type Plus struct {
	l Node
	r Node
}

func NewPlus(l, r Node) Node {
	return Plus{
		l,
		r,
	}
}

func (p Plus) Eval() float64 {
	return p.l.Eval() + p.r.Eval()
}
