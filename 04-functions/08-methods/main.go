package main

import (
	"fmt"
)

type Person struct {
	Fname  string
	Lname  string
	Age    uint8
	Salary float32
}

// FullName is a method that returns the full name of the
// person passed to it as a receiver
func (p Person) FullName() string {
	return p.Fname + " " + p.Lname
}

// GrowOld function makes the person grow older by 1 year
// and increases their salary
func (p Person) GrowOld() Person {
	p.Age += 1
	p.Salary = 1.1 * p.Salary
	return p
}

// GrowOldPtr does the same thing as GrowOld but using
// pointers without having to re assign
func (p *Person) GrowOldPtr() {
	p.Age += 1
	p.Salary = 1.1 * p.Salary
}

func main() {
	fmt.Println("Methods and receivers in Golang")

	// Okay lets create a person
	alice := Person{
		Fname:  "Alice",
		Lname:  "Petterson",
		Age:    23,
		Salary: 60000,
	}
	fmt.Println("The person is", alice.FullName())
	fmt.Println(alice.FullName(), "is", alice.Age, "yrs old and has a salary", alice.Salary)

	alice = alice.GrowOld()
	fmt.Println("1 year passes...")
	fmt.Println(alice.FullName(), "is", alice.Age, "yrs old and has a salary", alice.Salary)

	alice.GrowOldPtr()
	fmt.Println("1 year passes...")
	fmt.Println(alice.FullName(), "is", alice.Age, "yrs old and has a salary", alice.Salary)

}
