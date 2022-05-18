package main

import (
	"testing"
	"time"
)

var p Person = Person{
	Fname: "John",
	Lname: "Smith",
	Dob:   time.Date(1980, time.March, 21, 0, 0, 0, 0, time.UTC),
}

func TestPersonName(t *testing.T) {
	result := p.Name() // Get the output of the function
	er := "John Smith" // Specify the expected result

	// Check if the received output matches the expected result
	// or tell the test runner that the test failed
	if result != er {
		t.Error("Unexpected result: Was expecting", er, "but received", result)
	}
}

func TestPersonAge(t *testing.T) {
	r := p.Age()
	er := uint8(42)

	if r != er {
		t.Error("Unexpected result: Was expecting", er, "but received", r)
	}
}
