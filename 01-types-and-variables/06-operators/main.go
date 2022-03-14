package main

import "fmt"

func main() {
	fmt.Println("Operators in Golang")

	// Integer operators
	i1 := 10
	fmt.Println("Initial value of i1 =", i1)

	i2 := i1 + 15
	fmt.Println("Value of i1 + 15 =", i2)

	i2++
	fmt.Println("Value of i2 after increment =", i2)

	i2--
	fmt.Println("Value of i2 after decrement =", i2)

	i3 := i2 - 4
	fmt.Println("Value of i2 - 4 =", i3)

	i4 := i3 * i2
	fmt.Println("Value of i3 * i2 =", i4)

	i1 = i4 / 2
	fmt.Println("Value of i1 after reassignment =", i1)

	fmt.Println("Remainder of i4/2 =", i4%2)

	fmt.Println("Is i4 equal to i1 ?", i4 == i1)

	// Floating point operations
	f1 := 0.1
	fmt.Println("Value of f1 =", f1)

	f2 := f1 + 0.2
	fmt.Println("Value of f1 + 0.2 =", f2)

	f3 := float64(i3) - f2 // Type conversion int -> float64
	fmt.Println("Value of i3 + f2 =", f3)

	f1 = f2 * f3
	fmt.Println("Value of f2 * f3 =", f1)

	f4 := f1 / 3.14
	fmt.Println("Value of f1 / 3.14 =", f4)

	fmt.Println("Is f4 equal to f1 ?", f4 == f1)

	// Boolean logic operators
	hasBook := true
	hasLaptop := false

	fmt.Println("Do I have a book AND a laptop ?", hasBook && hasLaptop)
	fmt.Println("Do I have a book OR a laptop ?", hasBook || hasLaptop)
	fmt.Println("Do I not have a book ?", !hasBook)
	fmt.Println("Do I not have a laptop ?", !hasLaptop)
	fmt.Println("Do I not have a book AND not have a laptop ?", !hasBook && !hasLaptop)
	fmt.Println("Do I not have a book AND not have a laptop ?", !(hasBook || hasLaptop))

	num1 := 10
	num2 := 12

	fmt.Println("Is num1 > num2 ?", num1 > num2)
	fmt.Println("Is num1 < num2 ?", num1 < num2)
	fmt.Println("Is num1 not equal to num2 ?", num1 != num2)

	// String operators
	st1 := "Hello"
	fmt.Println("st1 =", st1)

	st2 := st1 + " " + "Golang" // String concatenation
	fmt.Println("st2 =", st2)

	st2 += "!!!" // add and re-assign
	fmt.Println("st2 after short assign =", st2)

	// Slice operations
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println("Original Slice", sl)

	sl1 := sl[10:15] // New slice with five elements from the middle
	fmt.Println("sl1 := sl[10:15] =", sl1)

	sl2 := sl[15:] // New slice from 15th index to last
	fmt.Println("sl2 := sl[15:] =", sl2)

	sl3 := sl[:10] // New slice from 1st element to just before 10th index
	fmt.Println("sl3 := sl[:10] =", sl3)

	sl4 := append(sl3, 21, 22, 23, 24, 25)
	fmt.Println("sl3 after appending values =", sl3)
	fmt.Println("sl4 after appending values to sl3 =", sl4)

	fmt.Println("Original sl =", sl)
	sl4 = append(sl[:11], sl[13:]...) // Removing values using append and unpacking operator
	fmt.Println("sl after removing index 11 and 12 =", sl4)

	//Bonus Check the length of slice
	fmt.Println("Length of sl4 =", len(sl4))
}
