package main

import "fmt"

type Person struct {
	Name   string
	Age    uint8
	Salary float32
	Pets   []Pet
}

type Pet struct {
	Name  string
	Age   uint8
	Cute  bool
	Bites bool
	Happy bool
}

func main() {
	fmt.Println("OOPS is POOPS. Structs FTW")

	p1 := Person{
		Name:   "Mr. Anderson",
		Age:    35,
		Salary: 100000,
	}

	fmt.Println("Here is p1", p1)

	pet1 := Pet{
		Name:  "Fluffy",
		Age:   5,
		Cute:  true,
		Bites: true,
		Happy: false,
	}

	p1.Pets = append(p1.Pets, pet1)
	fmt.Println("pets after adoption:", p1.Pets)

	var p2 Person
	p2.Age = 38
	p2.Name = "Agent Smith"
	p2.Salary = 0
	fmt.Println("Here is p2", p2)

	p2.Age++

	fmt.Println(p2.Name+"'s age after 1 year=", p2.Age)

	p3 := Person{
		Name:   "Trinity",
		Age:    33,
		Salary: 0,
		Pets: []Pet{
			{
				Name:  "Charlie",
				Age:   12,
				Cute:  false,
				Bites: true,
				Happy: true,
			},
			{
				Name:  "Tiger",
				Age:   6,
				Cute:  true,
				Bites: false,
				Happy: true,
			},
		},
	}
	fmt.Println("Person 3:", p3)
}
