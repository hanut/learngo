package main

import "fmt"

func main() {
	fmt.Println("Types and variables in Golang")
	var v1 bool = true
	var v2 int = -10
	var v3 float32 = 132.44
	var v4 uint = 123
	var v5 uint8 = 255
	var v6 string = "Hello Golang"

	fmt.Println("v1 =", v1)
	fmt.Println("v2 =", v2)
	fmt.Println("v3 =", v3)
	fmt.Println("v4 =", v4)
	fmt.Println("v5 =", v5)
	fmt.Println("v6 =", v6)

	// Null values in golang
	fmt.Println("Nil =", nil)
}
