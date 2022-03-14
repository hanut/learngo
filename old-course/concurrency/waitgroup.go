package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t", runtime.GOARCH)
	fmt.Println("CPUS:\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines:\t", runtime.NumGoroutine())
	fmt.Println("--------------------------------------")

	wg.Add(2)
	go foo()
	go bar()

	fmt.Println("CPUS:\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines:\t", runtime.NumGoroutine())
	fmt.Println("--------------------------------------")

	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("FOO:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("BAR:", i)
	}
	wg.Done()
}
