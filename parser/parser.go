package parser

import (
	"errors"
	"strconv"

	"github.com/NicoNex/calc/ops"
	"github.com/NicoNex/calc/utils"
)

// Type used to abstract the constructor functions of the operators.
type newOp func(l, r ops.Node) ops.Node

var precedence = map[string]int{
	"+": 0,
	"-": 0,
	"*": 1,
	"/": 1,
	"^": 2,
	"=": 3,
}

var InvalidOperator = errors.New("error: invalid operator")

func unwrap(i interface{}) item {
	return i.(item)
}

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
	case "^":
		return ops.NewPower, nil
	}
	return nil, InvalidOperator
}

// Converts a string operand to a float64 and returns it.
func parseOperand(o string) (float64, error) {
	return strconv.ParseFloat(o, 64)
}

// Returns the AST generated from the operators stack and operands queue.
func genAst(expr []item) ops.Node {
	var output = utils.NewStack()

	for i, itm := range expr {
		switch itm.typ {
		case itemOperand:
			val, err := parseOperand(itm.val)
			if err != nil {
				return nil
			}
			output.Push(ops.NewConst(val))

		case itemOperator:
			fn, err := parseOperator(itm.val)
			if err != nil {
				return nil
			}
			rnode := output.Pop()
			if rnode == nil {
				return nil
			}
			lnode := output.Pop()
			if lnode == nil {
				return nil
			}
			output.Push(fn(lnode.(ops.Node), rnode.(ops.Node)))

		case itemVariable:
			output.Push(ops.NewVariable(itm.val))

		case itemAssign:
			output.Pop()
			v := output.Pop()
			if v == nil {
				return nil
			}
			if i > 0 {
				expr[i] = expr[i-1]
				return ops.NewAssign(v.(ops.Variable), genAst(expr[i:]))
			}
		}
	}

	if ret := output.Pop(); ret != nil {
		return ret.(ops.Node)
	}
	return nil
}

// Returns true if a has precedence over b.
func hasPrecendence(a, b item) bool {
	return precedence[a.val] > precedence[b.val]
}

func toPostfix(items chan item) []item {
	var ret []item
	var stack = utils.NewStack()

	for i := range items {
		switch i.typ {
		case itemOperand, itemVariable:
			ret = append(ret, i)

		case itemOperator, itemAssign:
			for o := stack.Peek(); o != nil; o = stack.Peek() {
				if !hasPrecendence(unwrap(o), i) {
					break
				}
				ret = append(ret, o.(item))
				stack.Pop()
			}
			stack.Push(i)

		case itemBracket:
			switch i.val {
			case "(":
				stack.Push(i)
			case ")":
				for o := stack.Pop(); o != nil; o = stack.Pop() {
					if tmp := unwrap(o); tmp.val == "(" {
						break
					}
					ret = append(ret, o.(item))
				}
			}
		}
	}

	for o := stack.Pop(); o != nil; o = stack.Pop() {
		ret = append(ret, o.(item))
	}
	return ret
}

// Evaluates the types from the lexer and returns the AST.
func Parse(a string) ops.Node {
	_, items := lex(a)
	return genAst(toPostfix(items))
}
