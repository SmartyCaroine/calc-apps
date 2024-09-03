package handlers

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Calculator interface{ Calculate(a, b int) int }

type Handler struct {
	calculator Calculator
	output     io.Writer
}

func NewHandler(calculator Calculator, output io.Writer) *Handler {
	return &Handler{calculator: calculator, output: output}
}

func (this *Handler) handle(inputs []string) error {
	if len(inputs) < 2 {
		return fmt.Errorf("%w: two arguments required", errTooFewArguments)
	}

	a, err := strconv.Atoi(inputs[0])
	if err != nil {
		return fmt.Errorf("%w: input must be able to convert to integer", errMalformedInput)
	}
	b, err := strconv.Atoi(inputs[1])
	if err != nil {
		return fmt.Errorf("%w: input must be able to convert to integer", errMalformedInput)
	}

	c := this.calculator.Calculate(a, b)
	_, err = fmt.Fprintln(this.output, c)

	if err != nil {
		return errFailToWrite
	}

	return nil
}

var (
	errTooFewArguments = errors.New("Too few arguments were provided")
	errMalformedInput  = errors.New("One of the inputs was malformed")
	errFailToWrite     = errors.New("Output writer error")
)
