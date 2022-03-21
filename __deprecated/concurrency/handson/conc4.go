package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Fix the code in conc3.go using a mutex
func main() {

	fmt.Println("Number of CPUS", runtime.NumCPU())
	fmt.Println("Number of GoRoutines", runtime.NumGoroutine())

	const NUM_GO_ROUTINES = 10
	counter := 0
	mtx := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(NUM_GO_ROUTINES)

	for i := 0; i < NUM_GO_ROUTINES; i++ {
		go func() {
			mtx.Lock()
			tmp := counter
			runtime.Gosched()
			tmp++
			counter = tmp
			mtx.Unlock()
			wg.Done()
		}()
		fmt.Println("Active Routines:\t", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Final value of Counter:\t", counter)
}
