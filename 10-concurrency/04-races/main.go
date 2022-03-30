package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race conditions")
	counter := 0

	var wg sync.WaitGroup
	wg.Add(100)

	mtx := sync.Mutex{}

	for i := 0; i < 100; i++ {
		go func() {
			mtx.Lock()
			v := counter
			v++
			counter = v
			mtx.Unlock()
			wg.Done()
		}()
	}
	mtx.Lock()
	fmt.Println("Counter before wait completes", counter)
	mtx.Unlock()
	wg.Wait()

	fmt.Println("Counter:", counter)
}
