package main

import (
	"fmt"
	"strings"
)

var LINE = strings.Repeat("-", 80)

func Sep() {
	fmt.Println("\n" + LINE + "\n")
}

type Numeric interface {
	uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64
}

func SumOfNumbers[T Numeric](n []T) int {
	var s int
	for _, v := range n {
		s += int(v)
	}
	return s
}

func PrintMap[K comparable, V Numeric | bool](m map[K]V) {
	fmt.Printf("Our map is %v and has type %T\n", m, m)
	for k, v := range m {
		fmt.Println("Key=", k, "Value=", v)
	}
}

func main() {
	// Use SumOfNumbers() to get the sum of
	// different types of slices
	li := []int{1, 2, 3, 4, 5}
	li32 := []int32{1, 2, 3, 4, 5}
	lui8 := []uint8{1, 2, 3, 4, 5}
	lf64 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	sli := SumOfNumbers[int](li)
	sli32 := SumOfNumbers[int32](li32)
	slui8 := SumOfNumbers(lui8) // Works by type inference
	slf64 := SumOfNumbers(lf64) // Works by type inference

	fmt.Printf("Sum of %v (%T) = %d\n", li, li, sli)
	Sep()
	fmt.Printf("Sum of %v (%T) = %d\n", li32, li32, sli32)
	Sep()
	fmt.Printf("Sum of %v (%T) = %d\n", lui8, lui8, slui8)
	Sep()
	fmt.Printf("Sum of %v (%T) = %d\n", lf64, lf64, slf64)
	Sep()

	// Lets use the PrintMap()
	m1 := map[string]bool{"alice": true, "bob": false, "charlie": false}
	PrintMap(m1)
	m2 := map[int]float32{0: 10.1, 10: 20.2, 300: 33.33}
	PrintMap(m2)
	Sep()
}
