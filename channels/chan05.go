package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// Send values onto the channel
	go send(even, odd, quit)

	// Receive values
	receive(even, odd, quit)

}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v, ok := <-e:
			if !ok {
				continue
			}
			fmt.Println("Got an even value:", v)
		case v, ok := <-o:
			if !ok {
				continue
			}
			fmt.Println("Got an odd value:", v)
		case _, ok := <-q:
			if !ok {
				continue
			}
			fmt.Println("Quitting..")
			return
		}
	}
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- true
	close(q)
}
