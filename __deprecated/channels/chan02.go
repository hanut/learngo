package main

import "fmt"

func main() {
	cbuf := make(chan string, 1)
	cbuf <- "buggers"
	fmt.Println(<-cbuf)
}
