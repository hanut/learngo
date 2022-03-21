package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "$QWEasd@123#"
	fmt.Println("Original Password", s)

	bs, err := bcrypt.GenerateFromPassword([]byte(s), 10)

	if err != nil {
		fmt.Printf("Error:\t%v\n", err)
	}

	fmt.Println("Hashed Secret", string(bs))
}
