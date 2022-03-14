package main

import "fmt"

func main() {
	item := "necklace"

	// Simple switch statement
	switch item {
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
	s := true
	w := true

	switch {
	case s:
		fmt.Println("sunIsShining:", s) // Only this one is executed because its the first match
	case w:
		fmt.Println("Wont reach here") // Never reach here because switch exits after first match
	case s && w:
		fmt.Println("Wont reach here either") // Never going to come here either
	}

	// Same conditional switch with fallthru
	switch {
	case s:
		fmt.Println("The sun is shining.") // Only this one is executed because its the first match
		fallthrough
	case w:
		fmt.Println("The weather is sweet, yeah.")
		fallthrough
	case s && w:
		fmt.Println("Make you wanna move your dancing feet now.")
	}

	// Switch with initializer
	switch x := 3 + 4; {
	case x < 5:
		fmt.Println("Case One")
	case x == 5:
		fmt.Println("Case Two")
	case x > 5:
		fmt.Println("Case Three")
	}

	// Multiple Matches in same case
	role := "operator"

	switch role {
	case "manager", "operator":
		fmt.Println("No access to admin settings")
	case "admin":
		fmt.Println("Access to account admin")
	case "superadmin":
		fmt.Println("Full admin access")
	}
}
