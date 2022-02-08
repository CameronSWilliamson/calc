package main

import (
	"fmt"
	"os"
)

func main() {
	calc := newCalculator()
	result, err := calc.parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
}
