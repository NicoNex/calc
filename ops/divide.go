package ops

type Divide struct {
	l Node
	r Node
}

func NewDivide(l, r Node) Node {
	return Divide{
		l,
		r,
	}
}

func (d Divide) Eval() float64 {
	return d.l.Eval() / d.r.Eval()
}
