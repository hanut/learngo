package main

import "fmt"

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%d\t:\t%c\n", i, rune(i))
	}
}
