package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Printf("Outer Loop: %d\n", i)
		for j := 0; j < 4; j++ {
			fmt.Printf("\tInner Loop:%d", j)
		}
		fmt.Println()
	}
}
