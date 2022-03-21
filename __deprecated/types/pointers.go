package main

import "fmt"

func main() {
	x := 10
	y := &x
	fmt.Printf("Type: %T\tValue: %v\n", x, x)
	fmt.Printf("Type: %T\tValue: %v\n", y, y)

	Foo(x)
	fmt.Println("X outside Foo = ", x)

	Bar(&x)
	fmt.Println("X outside Bar = ", x)

	Bar(y)
	fmt.Println("X outside Bar = ", x)
}

func Foo(x int) {
	x++
	fmt.Println("X inside Foo = ", x)
}

func Bar(x *int) {
	(*x)++
	fmt.Println("X inside Bar = ", (*x))
}
