package main

import "fmt"

type Vehicle struct {
	Id     uint32
	Name   string
	Wheels uint8
}

type Car struct {
	AC          bool
	MusicSystem bool
	Wiper       bool
	Vehicle
}

type Scooter struct {
	KickStart bool
	Vehicle
}

func main() {
	fmt.Println("Reusing structs in Golang")

	car := Car{
		AC:          true,
		MusicSystem: true,
		Wiper:       true,
		Vehicle: Vehicle{
			Id:     123456,
			Name:   "Toyota Corolla",
			Wheels: 3,
		},
	}

	fmt.Println("Values of the Car:", car)

	// See how the properties are accessed from the reused struct
	fmt.Println("The name of the car is", car.Name)

	scooter := Scooter{KickStart: true}
	fmt.Println("Values of the Scooter:", scooter)

	// Lets set the name, id and Wheels of scooter
	scooter.Id = 12
	scooter.Name = "LML Vespa"
	scooter.Wheels = 2
	fmt.Println("Values of the Scooter:", scooter)
}
