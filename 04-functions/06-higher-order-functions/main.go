package main

import "fmt"

func main() {
	fmt.Println("Higher order Functions in Golang")

	// Simple HoF that builds greeting functions
	greetEnglish := GreetBuilder("Hello")
	greetHindi := GreetBuilder("Namaste")
	greetDeutsche := GreetBuilder("Guten Tag")
	greetEnglish("Alice")
	greetHindi("Alice")
	greetDeutsche("Alice")

	// Simple HoF that builds dynamic incrementor functions
	Inc3 := IncrementBuilder(3)
	Inc10 := IncrementBuilder(10)
	Inc100 := IncrementBuilder(100)

	x := 0
	fmt.Println("X:", x)
	x = Inc3(x)
	fmt.Println("X after Inc3():", x)
	x = Inc10(x)
	fmt.Println("X after Inc10():", x)
	x = Inc10(x)
	fmt.Println("X after Inc10():", x)
	x = Inc3(x)
	fmt.Println("X after Inc3():", x)
	x = Inc100(x)
	fmt.Println("X after Inc100():", x)
}

func GreetBuilder(greeting string) func(name string) {
	return func(name string) {
		fmt.Println(greeting + " ! " + name)
	}
}

// IncrementBuilder creates a new function that
// will increment given number according to its increment value
func IncrementBuilder(i int) func(n int) int {
	return func(n int) int {
		return n + i
	}
}
