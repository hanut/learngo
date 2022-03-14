package main

import "fmt"

func main() {
	fmt.Println("Pointers in Golang")

	// Simple integer variable
	x := 42
	fmt.Println("X =", x)

	// Simple pointer to x
	px := &x
	fmt.Println("Reference of X =", px)

	// Simple dereferencing
	y := *px
	fmt.Println("Dereferenced value of PX =", y)
}
