package main

import (
	"fmt"
)

func divide(a, b int) int {
	if b == 0 {
		panic("Division by zero is not allowed")
	}
	return a / b
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	numerator := 10
	denominator := 0
	fmt.Println("Not called here")
	result := divide(numerator, denominator)
	fmt.Println("called here")
	fmt.Println("Result:", result)
}
