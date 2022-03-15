package main

import "fmt"

func main() {
	fmt.Println("Returning values from a function")

	// Call the sum function and print the result
	sum := Sum(7, 8)
	fmt.Println("Sum of 7 + 8 is", sum)

	// Use a function to calculate the average value of an array of numbers
	nums := []int{120, 999, 540, 336, 771, 832, 986, 101, 255}
	avg := Average(nums)
	fmt.Println("Average of", nums, "=", avg)

	// Returning many values
	total, avg := TotalAndAvg(nums)
	fmt.Println("Total of", nums, "=", total, "and average is", avg)
}

// Simple function that returns a sum of two numbers
func Sum(a int, b int) int {
	return a + b
}

func Average(nums []int) float64 {
	var total float64
	for _, v := range nums {
		total += float64(v) // Convert to float64
	}
	return total / float64(len(nums)) // Convert len to float
}

func TotalAndAvg(nums []int) (float64, float64) {
	var total float64
	for _, v := range nums {
		total += float64(v) // Convert to float64
	}
	avg := total / float64(len(nums)) // Convert len to float
	return total, avg
}
