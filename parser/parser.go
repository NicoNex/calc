package parser

import (
	"os"
	"fmt"
	"strconv"

	"github.com/NicoNex/calc/ops"
	"github.com/NicoNex/calc/utils"
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

func parseOperand(o string) float64 {
	ret, err := strconv.ParseFloat(o, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return ret
}

func genAst(stack utils.Stack, queue utils.Queue) ops.Node {
	var ast ops.Node

	for operator, _ := stack.Pop(); operator != nil; {
		var node1 ops.Node
		var node2 ops.Node
		var tmp interface{}

		if ast == nil {
			tmp, err := queue.Pop()
			if err != nil {
				fmt.Println(err)
				return nil
			}
			node1 = tmp.(ops.Node)
		} else {
			node1 = ast
		}

		tmp, err := queue.Pop()
		if err != nil {
			fmt.Println(err)
			return nil
		}
		node2 = tmp.(ops.Node)
		opfn := operator.(newOp)
		ast = opfn(node1, node2)
	}

	return ast
}

func Parse(in string) ops.Node {
	var stack = utils.NewStack()
	var queue = utils.NewQueue()

	_, items := lex(in)

	for i := range items {
		switch i.typ {
		case operand:
			v := parseOperand(i.val)
			c := ops.NewConst(v)
			queue.Push(c)
		case operator:
			fn := parseOperator(i.val)
			stack.Push(fn)

		case variable:
			fmt.Println("variable")
		}
	}

	return genAst(stack, queue)
}
