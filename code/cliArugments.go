package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println(" ### CLI ARGUMENTS ###")

	var n1 = flag.Int("n1", 0, "Enter number 1")
	var n2 = flag.Int("n2", 0, "Enter number 2")
	flag.Parse()
	fmt.Println("n1: ", *n1)
	fmt.Println("n2: ", *n2)

	oout := *n1 + *n2
	fmt.Printf("Result: ")
	fmt.Println(oout)

	fmt.Printf("Try running the code with -h as CLI argument.")
}
