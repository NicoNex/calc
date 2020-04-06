package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/NicoNex/calc/parser"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>> ")
		string, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if ast := parser.Parse(string); ast != nil {
			fmt.Println(ast.Eval())
		} else {
			fmt.Println("syntax error")
		}
	}
}
