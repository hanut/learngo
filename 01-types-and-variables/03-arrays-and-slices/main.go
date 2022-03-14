package main

import "fmt"

func main() {
	fmt.Println("Arrays and Slices in Golang")
	// Simple array
	var ar1 [5]string
	ar1[0] = "Hello"
	ar1[1] = "And"
	ar1[2] = "Welcome"
	ar1[3] = "To"
	ar1[4] = "Golang"

	fmt.Println("Message in array", ar1)

	// Simple array with short declaration operation
	ar2 := [6]string{"Hello", "And", "Welcome", "Again", "To", "Golang"}
	fmt.Println("Message in array", ar2)

	// Simple slice
	sl1 := []string{"Once", "Again", "We", "Meet", "In", "Go"}
	fmt.Println("Message in slice", sl1)

	// appending to slices
	sl2 := append(sl1, "To", "Make", "Life", "Fun")
	fmt.Println("Message in new slice", sl2)

	// Accessing values by index
	fmt.Println("Word at index 3", sl2[3])

	// Strings are actually slices !!!
	var x []byte = []byte("Hello")
	fmt.Printf("The string as a byte slice =%v\tType of x =\t%T\n", x, x)
}
