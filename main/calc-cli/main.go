package main

import (
	"io"
	"os"

	"github.com/SmartyCaroine/calc-apps/handlers"

	"github.com/SmartyCaroine/calc-lib/calc"
)

func main() {
	var (
		inputs     []string            = os.Args[1:]
		calculator handlers.Calculator = calc.Addition{}
		output     io.Writer           = os.Stdout
	)
	handler := handlers.NewHandler(calculator, output)

	handler.Handle(inputs)
}
