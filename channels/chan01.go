package main

import "fmt"

func main() {
	c := make(chan string)
	go doStuff(c)
	// Print the value when it comes
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func doStuff(c chan string) {
	c <- "beans"
	c <- "bongs"
	c <- "thighs"
	c <- "thongs"
}
