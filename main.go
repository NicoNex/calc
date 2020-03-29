package main

import (
	"os"
	"fmt"
	"bufio"

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
		}
	}
}
