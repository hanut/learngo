package main

import "fmt"

type Person struct {
	Name string
	Age  uint8
}

func (p *Person) Speak() string {
	return fmt.Sprintf("Hi! My name is %s and I am %d years old.", p.Name, p.Age)
}

type Human interface {
	Speak() string
}

func saySomething(h Human) {
	fmt.Println(h.Speak())
}

func main() {
	people := []Person{}
	people = append(people, Person{Name: "Gunther Vaadim", Age: 22})
	people = append(people, Person{Name: "Vella Hoenheim", Age: 54})
	people = append(people, Person{Name: "Noel Hardrek", Age: 35})

	for _, p := range people {
		// saySomething(p)
		saySomething(&p)
	}
}
