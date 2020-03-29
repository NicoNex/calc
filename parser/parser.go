package parser

import (
	"os"
	"fmt"
	"strconv"

	"github.com/NicoNex/calc/ops"
	"github.com/NicoNex/calc/utils"
)

// Type used to abstract the constructor functions of the operators.
type newOp func(l, r ops.Node) ops.Node

// Parses the operator type and returns a function of type newOp
// according to the operator type.
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

// Converts a string operand to a float64 and returns it.
func parseOperand(o string) float64 {
	ret, err := strconv.ParseFloat(o, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return ret
}

// Returns the AST generated from the operators stack and operands queue.
func genAst(stack utils.Stack, queue utils.Queue) ops.Node {
	var ast ops.Node

	for opr, _ := stack.Pop(); opr != nil; opr, _ = stack.Pop() {
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
		opfn := opr.(newOp)
		ast = opfn(node1, node2)
	}

	return ast
}

// Evaluates the types from the lexer and returns the AST.
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
