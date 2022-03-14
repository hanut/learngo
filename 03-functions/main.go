package main

import "fmt"

func main() {
	nums := []int{4, 6, 8, 12, 20, 22, 123}
	x := Sum(nums...)
	fmt.Println("Result of Sum(4,6,8,12) =", x)
}

// Variadic Params
func Sum(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
