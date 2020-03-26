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

		fmt.Println(parser.Parse(string).Eval())
	}
}
