package main

import "fmt"

func main() {
	var v uint8

	// addition
	v = 5 + 1 + 10
	fmt.Println("Result of 5 + 1 + 10 =", v)

	// subtraction
	v = 10 - 5
	fmt.Println("Result of 10 - 5 =", v)

	// Increment
	v++
	fmt.Println("Result of ++ =", v)

	// Decrement
	v--
	fmt.Println("Result of -- =", v)

	// Decrement
	v = v * 2
	fmt.Println("Result of v*2 =", v)

	s := "Hello World"
	fmt.Println("Original String:", s)

	s = s + "! Lets take to the skies !!!"
	fmt.Println("After concatenation:", s)

	b := true
	fmt.Println("Original Bool:", b)

	b = !b
	fmt.Println("After (!):", b)
}
