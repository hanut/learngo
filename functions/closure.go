package main

import "fmt"

func main() {
	x := "beans"
	fmt.Printf("X = %v\ttypeof x = %T\n", x, x)

	{
		x := 123
		fmt.Printf("X = %v\t\ttypeof x = %T\n", x, x)
	}
	{
		x := false
		fmt.Printf("X = %v\ttypeof x = %T\n", x, x)
	}

	c := 0
	fmt.Println("C before increment is", c)

	func() {
		for i := 0; i < 10; i++ {
			c++
		}
	}()

	fmt.Println("C after increment is", c)

	c = 0
	fmt.Println("Before boost()", c)
	boost(&c, func() {
		c++
	})

	fmt.Println("After increment in callback func", c)

	y := makeLargeSlice(500000000)

	sum := int64(0)
	y.forEach(func(n int) {
		sum += int64(n)
	})
	fmt.Println("Sum of array", sum)
}

func boost(n *int, cb func()) {
	*n += 100
	fmt.Println("After boost()", *n)
	cb()
}

type LargeSlice []int

func makeLargeSlice(size int) LargeSlice {
	ms := make([]int, size, size)
	for i := 0; i < size; i++ {
		ms[i] = i + 1
	}
	return ms
}

func (s LargeSlice) forEach(f func(n int)) {
	for x := range s {
		f(x)
	}
}
