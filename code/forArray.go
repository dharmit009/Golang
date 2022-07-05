package main

import "fmt"

func main() {
	fmt.Println("For Array")
	var a = [3]string{"test", "testing", "tester"}

	for i, v := range a {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}
}
