// Write a program which prompts the user to enter a floating point number
// and prints the integer which is a truncated version of the floating point
// number that was entered. Truncation is the process of removing the digits
// to the right of the decimal place.

package main

import "fmt"

func main() {

	var number float32

	fmt.Printf("Enter a Floating number: ")
	fmt.Scanf("%f", &number)

	var number2 int32 = int32(number)
	fmt.Println("Truncated Number: ", number2)

}
