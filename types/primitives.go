package main

import (
	"fmt"
	"reflect"
)

func main() {
	// A simple integer type using short declaration operator
	v1 := 10
	fmt.Printf("Value:\t%v\t|\tType:\t%T\t|\tSize:\t%d bits\n", v1, v1, reflect.TypeOf(v1).Bits())

	// Simple integer using var & type for declaration
	var v2 int = 12345
	fmt.Printf("Value:\t%v\t|\tType:\t%T\t|\tSize:\t%d bits\n", v2, v2, reflect.TypeOf(v2).Bits())

	// Simple integer using var & type for declaration
	var v2_1 uint = 12345
	fmt.Printf("Value:\t%v\t|\tType:\t%T\t|\tSize:\t%d bits\n", v2_1, v2_1, reflect.TypeOf(v2_1).Bits())

	// Float64 using var & type for declaration
	var v3 float64 = 1200000000000000000
	fmt.Printf("Value:\t%v\t|\tType:\t%T\t|\tSize:\t%d bits\n", v3, v3, reflect.TypeOf(v3).Bits())

	// Float32 using var & type for declaration
	var v3_1 float32 = 90009
	fmt.Printf("Value:\t%v\t|\tType:\t%T\t|\tSize:\t%d bits\n", v3_1, v3_1, reflect.TypeOf(v3_1).Bits())

	// String using short declaration operator of string literal
	v4 := "hello"
	fmt.Printf("Value:\t%v\t|\tType:\t%T\t|\tSize:\t%d bytes\n", v4, v4, reflect.TypeOf(v4).Size())

	// boolean using short declaration operator of string literal
	v5 := false
	fmt.Printf("Value:\t%v\t|\tType:\t%T\t|\tSize:\t%d byte\n", v5, v5, reflect.TypeOf(v5).Size())

}
