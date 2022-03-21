package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Simple integers first
	var a int = 10
	fmt.Printf("Type: %T\tValue: %v\n", a, a)
	b := 10.1
	fmt.Printf("Type: %T\tValue: %v\n", b, b)
	var c uint8 = 255
	fmt.Printf("Type: %T\tValue: %v\n", c, c)

	// Parsing wrong type into a struct
	jstring := `{"value": 10000}`
	var decoded struct{ value uint8 }
	err := json.Unmarshal([]byte(jstring), &decoded)
	if err != nil {
		panic("Beep Boop error:" + err.Error())
	}
	fmt.Printf("Type: %T\tValue: %v\n", decoded, decoded)

	// Default Values or Zeroes
	var df1 int
	var df2 float64
	var df3 bool
	var df4 string
	var df5 uint
	var df6 []bool
	var df7 [5]bool
	df8 := make([]bool, 5)
	fmt.Printf("Type: %T\tValue: %v\n", df1, df1)
	fmt.Printf("Type: %T\tValue: %v\n", df2, df2)
	fmt.Printf("Type: %T\tValue: %v\n", df3, df3)
	fmt.Printf("Type: %T\tValue: %v\n", df4, df4)
	fmt.Printf("Type: %T\tValue: %v\n", df5, df5)
	fmt.Printf("Type: %T\tValue: %v\n", df6, df6)
	fmt.Printf("Type: %T\tValue: %v\n", df7, df7)
	fmt.Printf("Type: %T\tValue: %v\n", df8, df8)
}
