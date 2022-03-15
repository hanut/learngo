package main

import "fmt"

func main() {
	fmt.Println("Call by Value vs Call by Reference")

	// Simple example of how call by ref can mutate the value of the original variable
	x := 100

	byVal := IncrementByValue(x) // No escape so stack is used and this is faster
	fmt.Println("Value of x", x, "and value of result", byVal)

	byRef := IncrementByRef(&x) // This escapes and will cause the assigned space on heap so is slower
	fmt.Println("Value of x", x, "and value of result", byRef)

	//
}

// Uses a call by value approach
func IncrementByValue(x int) int {
	x = x + 1
	return x
}

// Uses a call by reference approach
func IncrementByRef(x *int) int {
	*x = *x + 1
	return *x
}
