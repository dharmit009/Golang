package main

import "fmt"

const normal, bold = "\033[0m", "\033[1m"

func main() {
	x := 0
	if x == 1 {
		fmt.Print("%v", x)
		fmt.Println("The value of x is 1")
	} else {
		fmt.Println("This is the else section")
	}
}
