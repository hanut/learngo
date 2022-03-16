package main

import (
	"fmt"
)

func main() {
	fmt.Println("Anonymous Functions in Golang")

	// A Simple anonymous ie. unnamed function
	myFunc := func(x int) int {
		return x + 5
	}
	y := 10
	r := myFunc(y)

	fmt.Println(y, "after myFunc =", r)

	// Javascript style, IFFE
	r = func(x int) int {
		return x + 12
	}(y)
	fmt.Println(y, "after IFFE =", r)

	// Passing an anonymous function as a callback
	Foo(y, func(x int) {
		fmt.Println(y, "after IFFE =", x)
	})
}

func Foo(x int, callback func(int)) {
	callback(x + 30)
}
