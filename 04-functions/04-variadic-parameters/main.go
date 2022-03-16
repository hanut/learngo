package main

import "fmt"

func main() {
	fmt.Println("Variadic paramters in Golang")

	// Calling a function that accepts variadic parameters to sum together values
	s := Sum(1, 2, 3, 4, 5)
	fmt.Println("Sum of 1,2,3,4,5 =", s)

	// Unpacking and array as variadic parameters to the sum function
	ns := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s = Sum(ns...)
	fmt.Println("Sum of 1,2,3,4,5,6,7,8,9,10 =", s)
}

// Sum is a function that takes any number of integer values as inputs and returns their Sum
func Sum(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}
