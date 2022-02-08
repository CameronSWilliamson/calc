package main

import (
	"strconv"
	"strings"
)


type calculator struct {
	Stack *stack
}

func newCalculator() *calculator {
	return &calculator{Stack: newStack()}
}

func (c *calculator) parse(calcCommands []string) (int, error) {
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
		default:
			floatVal, err := strconv.ParseFloat(calcCommand, 64)
			if err != nil {
				return 0, err
			}
			c.Stack.Push(int(floatVal))
		}
	}
	return c.Stack.Pop(), nil
}
