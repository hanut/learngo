package main

import "fmt"

type Numeric interface {
	uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64
}

// SumOfNumbers returns the sum of a slice of numbers (values that
// satisfy the Numeric interface)
// in the same type as that of the slice
func SumOfNumbers[T Numeric](n []T) T {
	var s T
	for _, v := range n {
		s += v
	}
	return s
}

func SumOfNumbersSpecificRt[T Numeric, R Numeric](n []T) R {
	var s R
	for _, v := range n {
		s += R(v)
	}
	return s
}

func main() {
	fmt.Println("Returning generic type values from a generic function")

	li := []int{1, 2, 3, 4, 5}
	lf64 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	lu8 := []uint8{1, 2, 3, 4, 5}

	s1 := SumOfNumbers(li)
	s2 := SumOfNumbers(lf64)
	s3 := SumOfNumbers(lu8)

	fmt.Printf("Sum of %v (%T) = %v (%T)\n", li, li, s1, s1)
	fmt.Printf("Sum of %v (%T) = %v (%T)\n", lf64, lf64, s2, s2)
	fmt.Printf("Sum of %v (%T) = %v (%T)\n", lu8, lu8, s3, s3)

	// Take a float64 slice of numbers and return an integer32 sum
	sint := SumOfNumbersSpecificRt[float64, int32](lf64)
	fmt.Printf("Sum with specific return of %v (%T) = %v (%T)\n", lf64, lf64, sint, sint)
	sint2 := SumOfNumbersSpecificRt[uint8, float64](lu8)
	fmt.Printf("Sum with specific return of %v (%T) = %v (%T)\n", lu8, lu8, sint2, sint2)
}
