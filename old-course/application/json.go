package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	p1 := Person{
		Id:   100001,
		Name: "Person One",
		Age:  22,
	}

	p2 := Person{
		Id:   100002,
		Name: "Person Two",
		Age:  22,
	}

	pj1, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	os.Stdout.Write(pj1)

	pj2, err := json.Marshal(p2)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	os.Stdout.Write([]byte(",\n"))
	os.Stdout.Write(pj2)
	os.Stdout.Write([]byte("\n\n"))

	jsonBlob := []byte(`[
		{"Name": "JP One", "Age": 20, "Id": 91000},
		{"Name": "JP Two", "Age": 21, "Id": 91001},
		{"Name": "JP Three", "Age": 22, "Id": 91002},
		{"Name": "JP Four", "Age": 22, "Id": 91003}
	]`)

	var people []Person
	err = json.Unmarshal(jsonBlob, &people)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	for _, person := range people {
		fmt.Printf("%s------------------------------------\n", person.String())
	}
}

type Person struct {
	Id   int64
	Name string
	Age  uint8
}

func (p *Person) String() string {
	return fmt.Sprintf("ID:\t%d\nNAME:\t%s\nAGE:\t%d\n", p.Id, p.Name, p.Age)
}
