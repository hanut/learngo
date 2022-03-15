package main

import "fmt"

const separator = "---------------------------------------"

func main() {
	// Structs are used to group logically related data
	// together
	fmt.Println("OOPS is POOPS. Structs FTW")

	// This struct holds status flags for a single conveyor belt
	type StatusFlags struct {
		isReady      bool
		isMoving     bool
		isJammed     bool
		isOverloaded bool
	}

	sf := StatusFlags{}

	fmt.Println("Single statusflag variable:", sf)

	// Imagine a bunch of conveyor belts and their statuses
	manysf := [3]StatusFlags{} // We have 3 conveyor belts now
	fmt.Println("Flags of all conveyors:", manysf)

	for idx, s := range manysf {
		fmt.Printf("Flags for Conveyor %d : %v\n", idx+1, s)
	}

	// Setting up the values of a struct
	sf2 := StatusFlags{
		isReady:      true,
		isMoving:     true,
		isJammed:     false,
		isOverloaded: false,
	}
	fmt.Printf("%s\nStatus of conveyor 2\n%s\n%v\n", separator, separator, sf2)

	// Access the properties of our struct to print a more readable message
	fmt.Printf("%s\nStatus of the conveyor 2\n%s\nReady:\t%t\nMoving:\t%t\nJammed:\t%t\nOverloaded:\t%t\n", separator, separator, sf2.isReady, sf2.isMoving, sf2.isJammed, sf2.isOverloaded)

	// We will now change the status to reflect that conveyor 2 is jammed
	sf2.isJammed = true
	fmt.Printf("%s\nStatus of the conveyor 2\n%s\nReady:\t%t\nMoving:\t%t\nJammed:\t%t\nOverloaded:\t%t\n", separator, separator, sf2.isReady, sf2.isMoving, sf2.isJammed, sf2.isOverloaded)

	// Lets update the value of isReady conveyor three in our array of conveyors
	manysf[2].isReady = true
	fmt.Printf("%s\nStatus of the conveyor 3\n%s\nReady:\t%t\nMoving:\t%t\nJammed:\t%t\nOverloaded:\t%t\n", separator, separator, manysf[2].isReady, manysf[2].isMoving, manysf[2].isJammed, manysf[2].isOverloaded)

}
