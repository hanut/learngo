package main

import "fmt"

func main() {
	y := "necklace"

	// Simple switch statement
	switch y {
	case "necklace":
		fmt.Println("Its a necklace")
	case "chain":
		fmt.Println("Its a chain")
	case "ring":
		fmt.Println("Its a ring")
	case "finding":
		fmt.Println("Its a finding")
	default:
		fmt.Println("I dont know what that is")
	}

	// Switch with conditions
	sunIsShining := true
	weatherIsSweet := true

	switch {
	case sunIsShining:
		fmt.Println("sunIsShining:", sunIsShining) // Only this one is executed because its the first match
	case weatherIsSweet:
		fmt.Println("Wont reach here") // Never reach here because switch exits after first match
	case sunIsShining && weatherIsSweet:
		fmt.Println("Wont reach here either") // Never going to come here either
	}

	// Same conditional switch with fallthru
	switch {
	case sunIsShining:
		fmt.Println("The sun is shining.") // Only this one is executed because its the first match
		fallthrough
	case weatherIsSweet:
		fmt.Println("The weather is sweet, yeah.")
		fallthrough
	case sunIsShining && weatherIsSweet:
		fmt.Println("Make you wanna move your dancing feet now.")
	}
}
