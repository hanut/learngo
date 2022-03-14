package main

import "fmt"

func main() {
	fmt.Println("Printing values and types in Golang")

	v1 := 42
	v2 := true
	v3 := "Hello Golang"

	// Formatted printing for formatting goodness
	fmt.Printf("Value of v1 =\t%d\nValue of v2 =\t%t\nValue of v3 =\t%s\n", v1, v2, v3)

	// Using %v for values
	fmt.Printf("Value of v1 =\t%v\nValue of v2 =\t%v\nValue of v3 =\t%v\n", v1, v2, v3)

	// Using %T for Types
	fmt.Printf("Type of v1 =\t%T\nType of v2 =\t%T\nType of v3 =\t%T\n", v1, v2, v3)
}
