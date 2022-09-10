package main

import "fmt"

func main() {

	fmt.Println("tagless-switch")
	var x = 3
	switch {
	case x < 5:
		fmt.Println("less than 5")

	case x > 5:
		fmt.Println("less than 5")
	}
}
