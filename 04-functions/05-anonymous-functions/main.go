package main

import (
	"fmt"
)

func main() {
	fmt.Println("Anonymous Functions in Golang")

	// A Simple anonymous ie. unnamed function
	myFunc := func(x int) int {
		return x + 5
	}
	y := 10
	r := myFunc(y)

	fmt.Println(y, "after myFunc =", r)

	// Javascript style, IIFE
	r = func(x int) int {
		return x + 12
	}(y)

	fmt.Println(y, "after IIFE =", r)

	// Passing an anonymous function as a callback
	IncCallback(y, func(x int) {
		fmt.Println(y, "after callback =", x)
	})

	// Callback function with named function argument
	IncCallback(y, PrintFunky)
}

func IncCallback(x int, callback func(int)) {
	iv := x + 20
	callback(iv)
}

// PrintFunky prints the number in a funky way
func PrintFunky(n int) {
	fmt.Println("Funky PrInTIng", n)
}
