// Package docsample provides a simple sum function for adding
// any number of argument integers
package docsample

// Adds stuff up and resturns the result
func Sum(nums ...int) int {
	sum := 0
	for n := range nums {
		sum += n
	}
	return sum
}
