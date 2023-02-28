package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	fmt.Println("### Random Number Generator ###")
	myarray := make([]byte, 10)
	fmt.Println("Before Generating Random Number: ")
	fmt.Println("Array: ", myarray)
	io.ReadFull(rand.Reader, myarray)
	fmt.Println("After Generating Random Number: ")
	fmt.Println("Array: ", myarray)
}
