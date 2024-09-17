package main

import (
	"flag"
	"os"

	"github.com/SmartyCaroine/calc-apps/handlers"

	"github.com/SmartyCaroine/calc-lib/calc"
)

func main() {
	var op string
	flag.StringVar(&op, "op", "+", "one of + - * /")
	flag.Parse()

	handler := handlers.NewHandler(calculators[op], os.Stdout)
	err := handler.Handle(flag.Args())

	if err != nil {
		panic(err)
	}
}

var calculators = map[string]handlers.Calculator{
	"+": calc.Addition{},
	"-": calc.Subtraction{},
	"/": calc.Division{},
	"*": calc.Multiplication{},
}
