package main

import "fmt"

func main() {
	names := []string{"Mark", "Tony", "Bittu", "Gajendra", "Carl"}
	fmt.Println("\nIteration using simple for loop")
	for i := 0; i < len(names); i++ {
		fmt.Println(i+1, names[i])
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
