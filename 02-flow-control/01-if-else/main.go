package main

import "fmt"

type Person struct {
	Name   string
	Age    uint8
	Salary float32
}

func main() {
	fmt.Println("Checking out if-else")

	rating := 7

	// If the rating is less than 5, let the person know it isn't good
	// otherwise tell them its awesome
	if rating < 5 {
		fmt.Printf("A rating of %d is not that great:\n", rating)
	} else {
		fmt.Printf("%d is a superb rating !\n", rating)
	}

	pi := 3.14

	// If pi == 1, everything is messed up, otherwise
	// if pi is still correct then the universe will survive
	// For all other cases the universe ends
	if pi == 1 {
		fmt.Println("We were wrong all along !")
	} else if pi == 3.14 {
		fmt.Println("The universe will survive")
	} else {
		fmt.Println("The universe fades to black...")
	}

	// Nested if-else
	// Generally try to avoid nested if else blocks as they lead
	// to unreadable code
	weight := 56
	if weight < 50 {
		fmt.Println("Too light.")
	} else {
		if weight < 70 {
			fmt.Println("Decent weight")
		} else {
			fmt.Println("Too heavy")
		}
	}

	// Simple if on a boolean value
	isReady := true

	if isReady {
		fmt.Println("Something is ready !!!")
	} else {
		fmt.Println("Something is not ready !!!")
	}

	// If with multiple conditions
	p1 := Person{
		Name:   "Alice",
		Salary: 90000,
		Age:    32,
	}

	if p1.Age > 30 && p1.Salary < 100000 {
		fmt.Println(p1.Name + " can do better")
	} else {
		fmt.Println(p1.Name + " is doing amazing")
	}
}
