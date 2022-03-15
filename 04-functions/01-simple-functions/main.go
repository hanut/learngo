package main

import "fmt"

func main() {
	fmt.Println("Playing with functions")
	Separate()

	// Function Call
	helloWorld()
	Separate()
	// Call the function with fixed parameters
	PrintSum(2, 3)
	Separate()
	PrintSum(12332, 86751)
	Separate()
	CombineStrings("Hello", "Alice")

}

func Separate() {
	fmt.Println("##################################")
}

// Function definition
func helloWorld() {
	fmt.Println("Hello World !")
}

// Function with fixed Parameters
func PrintSum(a int, b int) {
	fmt.Println("The sum of", a, "+", b, "=", a+b)
}

// FUnction to print combined strings
func CombineStrings(a string, b string) {
	fmt.Println("Combined String is:", a+" "+b)
}
