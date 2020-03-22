package ops

type Times struct {
	l Leaf
	r Leaf
}

func NewTimes(l, r Leaf) Times {
	return Times{
		l,
		r,
	}
}

func (t Times) Eval() float64 {
	return t.l.Eval() * t.r.Eval()
}
