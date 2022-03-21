package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUS:", runtime.NumCPU())
	fmt.Println("ROUTINES:", runtime.NumGoroutine())

	counter := int64(0)
	const gs = 100
	wg := sync.WaitGroup{}

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter:\t", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("ROUTINES:", runtime.NumGoroutine())
	}
	fmt.Println("ROUTINES:", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("Counter:", counter)
	fmt.Println("Done")
}
