package main

import "fmt"

func main() {
	x := 10
	fmt.Println("x =", x)

	{
		x++
		fmt.Println("x =", x)
	}
	fmt.Println("x =", x)

	{
		x := 20
		x++
		fmt.Println("x =", x)
	}

	fmt.Println("x =", x)
}
