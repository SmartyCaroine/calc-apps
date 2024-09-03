package calc_cli

import (
	"io"
	"os"
)

func main() {
	var (
		inputs     []string   = os.Args[1:]
		calculator Calculator = calc.Addition{}
		output     io.Writer  = os.Stdout
	)
	handler := NewHandler(calculator, output)

	handler.handle(inputs)
}
