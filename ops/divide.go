package ops

type Divide struct {
	l Leaf
	r Leaf
}

func NewDivide(l, r Leaf) Divide {
	return Divide{
		l,
		r,
	}
}

func (d Divide) Eval() float64 {
	return d.l.Eval() / d.r.Eval()
}
