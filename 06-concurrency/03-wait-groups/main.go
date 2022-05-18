package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ns := make([]uint32, 30000, 30000)

	for i := uint32(1); i <= 30000; i++ {
		ns[i-1] = i
	}
	s := SumOfSlice(ns)
	fmt.Println("The total sum is", s)

	var s2 uint32

	wg.Add(10) // We are running 10 async tasks together

	go SumOfSliceAsync(ns[:3000], &s2)
	go SumOfSliceAsync(ns[3000:6000], &s2)
	go SumOfSliceAsync(ns[6000:9000], &s2)
	go SumOfSliceAsync(ns[9000:12000], &s2)
	go SumOfSliceAsync(ns[12000:15000], &s2)
	go SumOfSliceAsync(ns[15000:18000], &s2)
	go SumOfSliceAsync(ns[18000:21000], &s2)
	go SumOfSliceAsync(ns[21000:24000], &s2)
	go SumOfSliceAsync(ns[24000:27000], &s2)
	go SumOfSliceAsync(ns[27000:], &s2)

	wg.Wait()

	fmt.Println("The async sum is", s2)
}

func SumOfSlice(s []uint32) uint32 {
	var sum uint32
	for _, v := range s {
		sum += v
	}
	return sum
}

func SumOfSliceAsync(s []uint32, sum *uint32) {
	ls := SumOfSlice(s)
	fmt.Println("Local sum", ls, "Active routines", runtime.NumGoroutine())
	*sum += ls
	wg.Done()
}
