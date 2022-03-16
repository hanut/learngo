package main

import "fmt"

func main() {
	fmt.Println("Variadic paramters in Golang")

	// Calling a function that accepts variadic parameters to sum together values
	s := Sum(1, 2, 3, 4, 5)
	fmt.Println("Sum of 1,2,3,4,5 =", s)

	// Unpacking and array as variadic parameters to the sum function
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s = Sum(nums...)
	fmt.Println("Sum of 1,2,3,4,5,6,7,8,9,10 =", s)

	// Variadic parameters must always be the last parameter
	sts := []string{"Alice", "Bob", "Charlie", "Dennis"}
	fmt.Println("Result of joining string slice:", JoinStrings('#', sts...))
}

// Sum is a function that takes any number of integer values as inputs and returns their Sum
func Sum(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func JoinStrings(j rune, s ...string) string {
	var t string

	for _, v := range s {
		t += v + string(j) // Type conversion of rune to string
	}

	return t
}
