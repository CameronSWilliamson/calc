package main

import (
	"math"
	"strconv"
	"strings"
)


type calculator struct {
	Stack *stack
}

func newCalculator() *calculator {
	return &calculator{Stack: newStack()}
}

func (c *calculator) parseRPN(calcCommands []string) (float64, error) {
	for _, calcCommand := range calcCommands {
		switch strings.ToLower(calcCommand) {
		case "+":
			c.Stack.Push(c.Stack.Pop() + c.Stack.Pop())
		case "-":
			c.Stack.Push(-c.Stack.Pop() + c.Stack.Pop())
		case "*", "x":
			c.Stack.Push(c.Stack.Pop() * c.Stack.Pop())
		case "/":
			c.Stack.Push(c.Stack.Pop() / c.Stack.Pop())
		case "sqrt":
			c.Stack.Push(math.Sqrt(c.Stack.Pop()))
		case "pi":
			c.Stack.Push(math.Pi)
		case "e":
			c.Stack.Push(math.E)
		default:
			floatVal, err := strconv.ParseFloat(calcCommand, 32)
			if err != nil {
				return 0, err
			}
			c.Stack.Push(floatVal)
		}
	}
	return c.Stack.Pop(), nil
}
