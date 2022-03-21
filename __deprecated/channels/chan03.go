package main

import "fmt"

var x uint8 = 0

func main() {
	c := make(chan uint8, 2)

	go doStuff(c)

	fmt.Println(<-c)
	fmt.Println(<-c)
}

func doStuff(c chan uint8) {
	x++
	c <- x
	x++
	c <- x
}
