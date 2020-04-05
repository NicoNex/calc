 package parser

import (
	"fmt"
	"errors"
	"strconv"

	"github.com/NicoNex/calc/ops"
	"github.com/NicoNex/calc/utils"
)

// Type used to abstract the constructor functions of the operators.
type newOp func(l, r ops.Node) ops.Node

// Parses the operator type and returns a function of type newOp
// according to the operator type.
func parseOperator(o string) (newOp, error) {
	switch o {
	case "+":
		return ops.NewPlus, nil
	case "-":
		return ops.NewMinus, nil
	case "*":
		return ops.NewTimes, nil
	case "/":
		return ops.NewDivide, nil
	}

	return nil, errors.New("error: invalid operator")
}

// Converts a string operand to a float64 and returns it.
func parseOperand(o string) (float64, error) {
	return strconv.ParseFloat(o, 64)
}

// Returns the AST generated from the operators stack and operands queue.
func genAst(expr utils.Queue) ops.Node {
	var output = utils.NewStack()

	for o, _ := expr.Pop(); o != nil; o, _ = expr.Pop() {
		var tmp = o.(item)

		switch tmp.typ {
		case operand:
			val, err := parseOperand(tmp.val)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			output.Push(ops.NewConst(val))

		case operator:
			fn, err := parseOperator(tmp.val)
			if err != nil {
				fmt.Println(err)
				return nil
			}

			rnode, err := output.Pop()
			if err != nil {
				fmt.Println(err)
				return nil
			}

			lnode, err := output.Pop()
			if err != nil {
				fmt.Println(err)
				return nil
			}

			output.Push(fn(rnode.(ops.Node), lnode.(ops.Node)))
		}
	}

	ret, err := output.Pop()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return ret.(ops.Node)
}

// Returns true if a has precedence over b.
func hasPrecendence(a, b item) bool {
	switch a.val {
	case "+", "-":
		return false

	case "*", "/":
		return b.val != "*" && b.val != "/"
	}
	return false
}

// Evaluates the types from the lexer and returns the AST.
func Parse (a string) ops.Node {
	var stack = utils.NewStack()
	var queue = utils.NewQueue()

	_, items := lex(a)

	for i := range items {
		switch i.typ {
		case operand:
			queue.Push(i)

		case operator:
			for o, _ := stack.Peek(); o != nil; o, _ = stack.Peek() {
				var tmp = o.(item)

				if hasPrecendence(tmp, i) {
					queue.Push(o)
					stack.Pop()
				} else {
					break
				}
			}
			// TODO: handle brackets here
			stack.Push(i)

		case variable:
			fmt.Println("variable")
		}
	}

	for o, _ := stack.Pop(); o != nil; o, _ = stack.Pop() {
		queue.Push(o)
	}

	return genAst(queue)
}
