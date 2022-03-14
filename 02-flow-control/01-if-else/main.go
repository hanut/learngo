package main

import "fmt"

type Person struct {
	Name   string
	Age    uint8
	Salary float32
}

func main() {
	rating := 7

	if rating < 5 {
		fmt.Printf("A rating of %d is not that great:\n", rating)
	} else {
		fmt.Printf("%d is a superb rating !\n", rating)
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
