package main

import (
	"fmt"
	"time"
)

const separator = "---------------------------------------"

type StatusFlags struct {
	isReady      bool
	isMoving     bool
	isJammed     bool
	isOverloaded bool
}

type Conveyor struct {
	Id        uint8
	Status    StatusFlags
	StartedAt time.Time
}

func main() {
	fmt.Println("Nested structs in Golang")

	// Lets create a new conveyor
	cv := Conveyor{
		Id:        1,
		StartedAt: time.Now(),
		Status: StatusFlags{
			isReady: true,
		},
	}

	fmt.Printf("%s\nStatus of the Conveyor(%d)\n%s\nReady:\t%t\nMoving:\t%t\nJammed:\t%t\nOverloaded:\t%t\n", separator, cv.Id, separator, cv.Status.isReady, cv.Status.isMoving, cv.Status.isJammed, cv.Status.isOverloaded)

	// Lets update the status of the conveyor to show its jammed
	cv.Status.isJammed = true
	fmt.Printf("%s\nStatus of the Conveyor(%d)\n%s\nReady:\t%t\nMoving:\t%t\nJammed:\t%t\nOverloaded:\t%t\n", separator, cv.Id, separator, cv.Status.isReady, cv.Status.isMoving, cv.Status.isJammed, cv.Status.isOverloaded)

	// Lets show the user when the Conveyor started if it is ready
	if cv.Status.isReady {
		fmt.Println("The conveyor started at:", cv.StartedAt)
	}
}
