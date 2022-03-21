package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	io.WriteString(os.Stdout, "Hello World\n")
	fmt.Println(os.TempDir())
}
