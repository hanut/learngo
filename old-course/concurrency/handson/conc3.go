package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Create a program with a race condition
// and prove it has a race condition by using
// go run -race to check for races. Use wait groups to
// pause execution
func main() {

	fmt.Println("Number of CPUS", runtime.NumCPU())
	fmt.Println("Number of GoRoutines", runtime.NumGoroutine())

	const NUM_GO_ROUTINES = 10
	counter := 0

	wg := sync.WaitGroup{}
	wg.Add(NUM_GO_ROUTINES)

	for i := 0; i < NUM_GO_ROUTINES; i++ {
		go func() {
			tmp := counter
			runtime.Gosched()
			tmp++
			counter = tmp
			wg.Done()
		}()
		fmt.Println("Active Routines:\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Final value of Counter:\t", counter)
}
