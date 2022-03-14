package main

import "fmt"

func main() {
	helloWorld()
	greetPerson("Hanut")
	greetManyPeople("George", "Matthew", "Akhilesh", "Obumbe", "Alexandria")
	people := []string{"Aditya", "Chirag", "Tarun", "Chacha", "Anu"}
	greetManyPeople(people...)

	s := sum(10000, 4000231)
	fmt.Println("Sum of 10000 + 4000231 = ", s)

	nums := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := sumMany(nums...)
	fmt.Println("Sum of", nums, " = ", s2)
}

func helloWorld() {
	fmt.Println("Hello World")
}

func greetPerson(name string) {
	fmt.Println("Hello", name)
}

func greetManyPeople(people ...string) {
	for _, person := range people {
		fmt.Printf("Hello %s.\n", person)
	}
}

func sum(a float64, b float64) float64 {
	return a + b
}

func sumMany(nums ...float64) float64 {
	var sum float64 = 0

	for _, n := range nums {
		sum += n
	}
	return sum
}
