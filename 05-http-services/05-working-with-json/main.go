package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string
	Age    uint8
	Height uint8
	Weight uint8
	Status bool
}

type People []Person

func main() {
	fmt.Println("Working with json in Golang")

	lp := People{
		{Name: "Alice", Age: 22, Height: 144, Weight: 65, Status: true},
		{Name: "Bob", Age: 40, Height: 195, Weight: 120, Status: false},
		{Name: "Charlie", Age: 25, Height: 185, Weight: 80, Status: true},
		{Name: "Dan", Age: 32, Height: 188, Weight: 101, Status: false},
		{Name: "Edward", Age: 18, Height: 160, Weight: 74, Status: true},
	}

	fmt.Println("\nOriginal list of people", lp)

	// Marshalling (Encode) a struct to json string
	js, err := json.Marshal(lp)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nThe converted list in json is :", string(js)) // Type conversion from []byte -> string and print

	// Unmarshalling (Decode) json into a struct or variable
	var dl People
	err = json.Unmarshal(js, &dl)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nDecoded people list", dl)

	// // Converting a json string into a golang struct
	jsPerson := `{"Name":"Some Test User", "Age": 45, "Weight": 101, "Height": 180, "Status": false}`

	var someGuy Person
	err = json.Unmarshal([]byte(jsPerson), &someGuy)

	fmt.Println("\nThe new guy", someGuy)

}
