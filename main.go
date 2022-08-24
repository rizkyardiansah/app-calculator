package main

import (
	"fmt"
	go_calculator "github.com/rizkyardiansah/go-calculator"
)

func main() {
	fmt.Println(go_calculator.Add(1,2))
	fmt.Println(go_calculator.Substract(1,2))
	fmt.Println(go_calculator.Multiply(1,2))
	fmt.Println(go_calculator.Divide(1,2))
	fmt.Println(go_calculator.Modulo(1,2))
	fmt.Println(go_calculator.Power(1,2))
}