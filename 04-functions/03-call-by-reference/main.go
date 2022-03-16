package main

import "errors"

func main() {
	println("Call by Value vs Call by Reference")

	// Simple example of how call by ref can mutate the value of the original variable
	x := 100

	byVal := IncrementByValue(x) // No escape so stack is used and this is faster
	println("Value of x", x, "and value of result", byVal)
	byRef, _ := IncrementByRef(&x) // No escape but we would need to do nil pointer checks

	println("Value of x", x, "and value of result", byRef)

	withPointer, _ := ReturnAPointer(&x)
	println("Value of x", x, "and value of result", withPointer)

}

// Uses a call by value approach
func IncrementByValue(x int) int {
	x = x + 1
	return x
}

// Uses a call by reference approach
func IncrementByRef(x *int) (int, error) {
	if x == nil {
		return 0, errors.New("The passed pointer is nil")
	}
	*x = *x + 1
	return *x, nil
}

func ReturnAPointer(x *int) (*int, error) {
	if x == nil {
		return nil, errors.New("Received nil pointer in ReturnAPointer()")
	}
	*x = *x + 1
	return x, nil
}
