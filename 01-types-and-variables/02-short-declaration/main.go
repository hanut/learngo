package main

import "fmt"

func main() {
	fmt.Println("Variables using the short declaration operator")

	v1 := 10
	v2 := uint8(255) // Type conversion integer -> uint8
	v3 := true
	v4 := 10.04
	v5 := float32(10.04) // Type conversion float64 -> float32
	v6 := "Hello Golang"

	fmt.Println("v1 =", v1)
	fmt.Println("v2 =", v2)
	fmt.Println("v3 =", v3)
	fmt.Println("v4 =", v4)
	fmt.Println("v5 =", v5)
	fmt.Println("v6 =", v6)
}
