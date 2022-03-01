package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// Send values onto the channel
	go send(even, odd, quit)

	// Receive values
	receive(even, odd, quit)

}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Got an even value:", v)
		case v := <-o:
			fmt.Println("Got an odd value:", v)
		case <-q:
			fmt.Println("Quitting..")
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}
