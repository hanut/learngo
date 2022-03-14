package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// Fix the code in conc3.go using a atomic
func main() {

	fmt.Println("Number of CPUS", runtime.NumCPU())
	fmt.Println("Number of GoRoutines", runtime.NumGoroutine())

	const NUM_GO_ROUTINES = 10
	counter := int32(0)
	wg := sync.WaitGroup{}
	wg.Add(NUM_GO_ROUTINES)

	for i := 0; i < NUM_GO_ROUTINES; i++ {
		go func() {
			atomic.AddInt32(&counter, 1)
			runtime.Gosched()
			fmt.Println("EDITED VALUE:\t", atomic.LoadInt32(&counter))
			wg.Done()
		}()
		fmt.Println("Active Routines:\t", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Final value of Counter:\t", counter)
}
