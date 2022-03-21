package main

import "fmt"

func main() {
	add10 := adderHoC(10)
	fmt.Println("Adding 10 to 100", add10(100))

	add42 := adderHoC(42)
	fmt.Println("Adding 42 to 100", add42(100))
}

func adderHoC(val int) func(num int) int {
	return func(num int) int {
		return num + val
	}
}
