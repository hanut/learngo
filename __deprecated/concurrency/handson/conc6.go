package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Operating System:\t", runtime.GOOS)
	fmt.Println("Architecture:\t\t", runtime.GOARCH)
	fmt.Println("Execution Root:\t\t", runtime.GOROOT())
	fmt.Println("Available CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Go Version:\t\t", runtime.Version())
}
