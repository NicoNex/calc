package ops

type Times struct {
	l Node
	r Node
}

func NewTimes(l, r Node) Node {
	return Times{
		l,
		r,
	}
}

func (t Times) Eval() float64 {
	return t.l.Eval() * t.r.Eval()
}
