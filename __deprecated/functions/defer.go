package main

import "fmt"

func main() {
	defer Foo()
	Bar()
}

func Foo() {
	fmt.Println("Called Foo")
}

func Bar() {
	fmt.Println("Called Bar")
}
