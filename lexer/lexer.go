package lexer

import (
	"os"
	"fmt"
	"regexp"
	"strconv"

	"github.com/NicoNex/calc/ops"
)

type newop func(a, b ops.Node) ops.Node

const (
	operator = `\+|\-|\*|\/|\=`
	constant = `-?\d+(,\d+)*(\.\d+(e\d+)?)?`
	variable = `\w*`
)

func die(msg ...interface{}) {
	for _, m := range msg {
		fmt.Println(m)
	}
	os.Exit(1)
}

func check(e error) {
	if e != nil {
		die("syntax error")
	}
}

func split(s string) []string {
	re := regexp.MustCompile(`(-?\d+(,\d+)*(\.\d+(e\d+)?)?|\+|\-|\*|\/|\=|\w*)`)
	return re.FindAllString(s, -1)
}

func populateAst(toks []string, node *ops.Node, i int) ops.Node {
	return nil
}

func parseOperator(operator string) newop {
	switch operator {
	case "+":
		return ops.NewPlus
	case "-":
		return ops.NewMinus
	case "*":
		return ops.NewTimes
	case "/":
		return ops.NewDivide

	default:
		die("error: unknown operator")
	}

	return nil
}

func Parse(s string) ops.Node {
	var tmp ops.Node
	var tree ops.Node
	var tokens = split(s)
	var rop = regexp.MustCompile(operator)
	var rco = regexp.MustCompile(constant)
	// var rva = regexp.MustCompile(variable)

	for k, t := range tokens {
		switch {

		case rco.MatchString(t):
			n, err := strconv.ParseFloat(t, 64)
			check(err)
			tmp = ops.NewConst(n)

		case rop.MatchString(t):
			n, err := strconv.ParseFloat(tokens[k+1], 64)
			check(err)
			newOperator := parseOperator(t)

			if tree != nil {
				tree = newOperator(tree, ops.NewConst(n))
			} else {
				tree = newOperator(tmp, ops.NewConst(n))
			}
		}
	}


	return tree
}
