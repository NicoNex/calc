package main

import (
	"fmt"

	"github.com/NicoNex/calc/ops"
)

func main() {
	res := ops.NewTimes(
		ops.NewConst(3),
		ops.NewPlus(
			ops.NewConst(5),
			ops.NewConst(7),
		),
	)

	res2 := ops.NewPlus(
		ops.NewConst(2),
		ops.NewPlus(
			ops.NewConst(5),
			ops.NewConst(5),
		),
	)

	fmt.Println(res.Eval(), res2.Eval())
}
