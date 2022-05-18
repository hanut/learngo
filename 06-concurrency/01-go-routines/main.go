package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Go Routines!!!")

	// Get the number of available machine cores
	c := runtime.NumCPU()
	fmt.Println("Your machine has", c, "cores available")

	sa := runtime.GOARCH
	os := runtime.GOOS

	fmt.Printf("Your system is %s and runs %s\n", sa, os)

	ng := runtime.NumGoroutine()
	fmt.Println(ng, "go routine(s) are executing right now")

	func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Anon complete")
	}()

	fmt.Println(ng, "go routine(s) are executing right now")

	// This anonymous go routine will not print anything
	// if the  main() function doesnt sleep as
	// main will exit before it can complete
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Go routine complete")
	}()

	time.Sleep(time.Second * 3)
}
