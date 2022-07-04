package main

import "fmt"

func main() {
	var x int = 1
	var y int = 0

	var pointer *int = &x

	fmt.Println("Before using y=*p (pointer)")
	fmt.Println("X: ", x)
	fmt.Println("Y: ", y)

	y = *pointer

	fmt.Println("\nAfter using y=*p (pointer)")
	fmt.Println("X: ", x)
	fmt.Println("Y: ", y)

	ptr := new(int)
	*ptr = 4
	fmt.Println("\nValue Of ptr: ", *ptr)
	fmt.Println("Address Of ptr: ", ptr)
}
