package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go foo(c)
	go bar(c)
	baz(c)

	fmt.Println("Exiting...")
}

func foo(c chan<- string) {
	c <- "Hello from foo"
	time.Sleep(time.Second * 3)
	c <- "Delayed message from foo"
}

func bar(c <-chan string) {
	for msg := range c {
		fmt.Println("Message Received:", msg)
	}
}

func baz(c chan string) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("Baz is Baz (%d)", i+1)
	}
	fmt.Println("Received in Baz:", <-c)
	close(c)
}
