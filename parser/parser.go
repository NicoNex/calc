package parser

import (
	"fmt"
	"github.com/NicoNex/calc/ops"
)

type newOp func(l, r ops.Node) ops.Node

func parseOperator(o string) newOp {
	switch o {
	case "+":
		return ops.NewPlus
	case "-":
		return ops.NewMinus
	case "*":
		return ops.NewTimes
	case "/":
		return ops.NewDivide
	}

	return nil
}

func Parse(in string) ops.Node {
	// var tmp ops.Node
	// var prev item

	_, items := lex(in)

	for i := range items {
		switch i.typ {
		case constant:
			fmt.Println("constant")
			// prev = i
		case operator:
			fmt.Println("operator")

		case variable:
			fmt.Println("variable")

		}
	}

	return ops.NewConst(666)
}
