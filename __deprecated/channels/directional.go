package main

import "fmt"

func main() {
	cbidi := make(chan int)
	fmt.Printf("%T\n", cbidi)

	csend := make(chan<- int)
	fmt.Printf("%T\n", csend)

	crec := make(<-chan int)
	fmt.Printf("%T\n", crec)
}
