package main

import (
	"fmt"
	"time"
)

type Person struct {
	Fname string
	Lname string
	Dob   time.Time
}

func (p Person) Name() string {
	return p.Fname + " " + p.Lname
}

func (p Person) Age() uint8 {
	hrs := time.Since(p.Dob).Hours()
	yrs := hrs / (24 * 365)
	return uint8(yrs)
}

// This is just here for show
func main() {
	p := Person{
		Fname: "John",
		Lname: "Smith",
		Dob:   time.Date(1980, 2, 10, 0, 0, 0, 0, time.UTC),
	}

	fmt.Println(p)
	fmt.Println(p.Age())
	fmt.Println(p.Name())
}
