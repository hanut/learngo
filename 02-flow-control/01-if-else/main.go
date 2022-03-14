package main

import "fmt"

func main() {
	rating := 7

	if rating < 5 {
		fmt.Println("Not that great of a rating:", rating)
	} else {
		fmt.Println("Superb rating :", rating)
	}

	// Simple if on a boolean value
	hasBeans := true

	if hasBeans {
		fmt.Println("My man has beans !!!")
	}

}
