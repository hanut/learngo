package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go foo(&wg)
	go bar(&wg)

	wg.Wait()
}

func foo(wg *sync.WaitGroup) {
	fmt.Println("Hello from Foo()")
	wg.Done()
}

func bar(wg *sync.WaitGroup) {
	fmt.Println("Hello from Bar()")
	wg.Done()
}
