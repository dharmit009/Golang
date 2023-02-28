package main

import "fmt"

func average(a, b, c int) int {
	return (a + b + c) / 3
}

func main() {
	fmt.Println("### If Else-If stateement")
	var (
		a int = 1
		b int = 2
		c int = 3
	)

	if a > b {
		fmt.Println("A is greater than B")
	} else if b > c {
		fmt.Println("B is greater than C")
	} else {
		fmt.Println("C is greater than A and B")
	}

	if average(a, b, c) > 4 {
		fmt.Println("Acceptable")
	} else {
		fmt.Println("Unacceptable")
	}
}
