package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		help()
	}
	calc := newCalculator()
	result, err := calc.parseRPN(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%f\n", result)
}

func help() {
	fmt.Println("Reverse Polish Notation Calculator")
	fmt.Println("Usage: calc <expression>")
	fmt.Println("")
	fmt.Println("Example:")
	fmt.Println("\tcalc 6 6 *")
	fmt.Println("\t> 36")
	fmt.Println("")
	fmt.Println("The following operators are supported:")
	fmt.Println("+ - * / sqrt pi e")
	os.Exit(1)
}