package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	fmt.Println("CPUS:", runtime.NumCPU())
	fmt.Println("ROUTINES:", runtime.NumGoroutine())

	const gs = 100
	wg := sync.WaitGroup{}

	mtx := sync.Mutex{}

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mtx.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mtx.Unlock()
			wg.Done()
		}()
	}
	fmt.Println("ROUTINES:", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("Counter:", counter)
	fmt.Println("Done")
}
