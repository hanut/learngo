package main

import "fmt"

func main() {
	y := "necklace"

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

}
