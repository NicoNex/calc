package parser

import (
<<<<<<< Updated upstream
	"os"
	"fmt"
=======
	"fmt"
	"errors"
>>>>>>> Stashed changes
	"strconv"

	"github.com/NicoNex/calc/ops"
	"github.com/NicoNex/calc/utils"
)

// Type used to abstract the constructor functions of the operators.
type newOp func(l, r ops.Node) ops.Node
<<<<<<< Updated upstream
=======

var precedence = map[string]int{
	"+": 0,
	"-": 0,
	"*": 1,
	"/": 1,
	"^": 2,
	"=": 3,
}

// Could be float64 but it's ops.Node for better scalability.
var dict map[string]ops.Node
>>>>>>> Stashed changes

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

func castToItem(o interface{}, err error) (item, error) {
	if it, ok := o.(item); ok {
		return it, err
	}
	return item{}, err
}

func castItem(i interface{}, e error) (item, error) {
	if e != nil {
		if it, ok := i.(item); ok {
			return it, nil
		}
		return item{}, errors.New("error: type assertion failed")
	}
	return item{}, e
}

// Returns the AST generated from the operators stack and operands queue.
func genAst(stack utils.Stack, queue utils.Queue) ops.Node {
	var ast ops.Node

<<<<<<< Updated upstream
	for opr, _ := stack.Pop(); opr != nil; opr, _ = stack.Pop() {
		var node1 ops.Node
		var node2 ops.Node
		var tmp interface{}

		if ast == nil {
			tmp, err := queue.Pop()
=======
	for itm, err := castItem(expr.Pop());
			err != nil; itm, err = castItem(expr.Pop()) {

		switch itm.typ {
		case operand:
			val, err := parseOperand(itm.val)
>>>>>>> Stashed changes
			if err != nil {
				fmt.Println(err)
				return nil
			}
<<<<<<< Updated upstream
			node1 = tmp.(ops.Node)
		} else {
			node1 = ast
=======
			output.Push(ops.NewConst(val))

		case operator:
			fn, err := parseOperator(itm.val)
			if err != nil {
				return nil
			}
			rnode, err := output.Pop()
			if err != nil {
				return nil
			}
			lnode, err := output.Pop()
			if err != nil {
				return nil
			}
			output.Push(fn(lnode.(ops.Node), rnode.(ops.Node)))

		case assign:


		case variable:
			next, err := castItem(expr.Peek())
			if err != nil {
				if error.Is(err, utils.EmptyQueue) {
					if node, ok := dict[itm.val]; ok {
						return node
					}
				}
				return nil
			}
			if val, ok := dict[itm.val]; ok && next.typ != assign {
				output.Push(val)
			}

		// 	switch next.typ {
		// 	case assign
		// 	}

			// TODO: can't continue without proper errors
			// check if next is assign and proceed...
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
		case operand:
			v := parseOperand(i.val)
			c := ops.NewConst(v)
			queue.Push(c)
		case operator:
			fn := parseOperator(i.val)
			stack.Push(fn)
=======
		case operand, variable:
			queue.Push(i)

		case operator, assign:
			for o, _ := stack.Peek(); o != nil; o, _ = stack.Peek() {
				if tmp := o.(item); !hasPrecendence(tmp, i) {
					break
				}
				queue.Push(o)
				stack.Pop()
			}
			stack.Push(i)

		case bracket:
			switch i.val {
			case "(":
				stack.Push(i)
			case ")":
				for o, _ := stack.Pop(); o != nil; o, _ = stack.Pop() {
					if tmp := o.(item); tmp.val == "(" {
						break
					}
					queue.Push(o)
				}
			}
		}
	}
>>>>>>> Stashed changes

		case variable:
			fmt.Println("variable")
		}
	}

<<<<<<< Updated upstream
	return genAst(stack, queue)
=======
	fmt.Println(queue)
	return genAst(queue)
>>>>>>> Stashed changes
}
