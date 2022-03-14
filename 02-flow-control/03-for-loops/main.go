package main

import "fmt"

func main() {
	fmt.Println("Looping in Golang")

	// Define an array of 5 string names
	names := [5]string{"Mark", "Tony", "Bittu", "Gajendra", "Carl"}

	fmt.Println("\nIteration using simple for loop")
	for i := 0; i < len(names); i++ {
		fmt.Printf("%d.\t%s\n", i+1, names[i])
	}

	fmt.Println("\nIteration using range")
	for index, name := range names {
		fmt.Printf("%d.\t%s\n", index+1, name)
	}

	fmt.Println("\nIteration using while loop style syntax")
	counter := 0
	for {
		fmt.Printf("%d.\t%s\n", counter+1, names[counter])
		counter++
		if counter > len(names)-1 {
			break
		}
	}
}
