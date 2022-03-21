package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous function")
	}()

	func(name string) {
		fmt.Println("Hello", name)
	}("Hanut Singh Gusain")

	foo(func(s string) {
		fmt.Println("Hello ", s)
	})

	bar := func(s string) {
		fmt.Println(s)
	}

	bar("boom boom bhati")

	baz := func(x func(s string)) {
		fmt.Println("Running the func expression")
		x("broderick")
	}

	baz(func(s string) { fmt.Println("received ", s, "in callback") })
}

func foo(cb func(s string)) {
	fmt.Println("Running foo")
	cb("beans")
}
