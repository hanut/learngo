package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("More go routines")

	go DelayedPrinter("Message One", time.Millisecond*300)
	go DelayedPrinter("Message Two", time.Millisecond*300)
	go DelayedPrinter("Message Three", time.Millisecond*300)
	go DelayedPrinter("Message Four", time.Millisecond*300)
	go DelayedPrinter("Message Five", time.Millisecond*300)

	time.Sleep(time.Millisecond * 100)
	fmt.Println("Active Go routines", runtime.NumGoroutine())
	time.Sleep(time.Second)
}

// DelayedPrinter uses fmt.Println to output any input
// v after a delay d
func DelayedPrinter(v interface{}, d time.Duration) {
	time.Sleep(d)
	fmt.Println(v)
}
