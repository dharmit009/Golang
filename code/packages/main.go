package main

import (
	"fmt"

	"example.com/mypack/greeter"
	"example.com/mypack/operations"
)

func main() {

	fmt.Println("### Packages ###")
	fmt.Println(greeter.Greeter("World!"))
	fmt.Println(greeter.Greeter("James!"))
	fmt.Println(greeter.Greeter("Testing!"))

	var out float32 = operations.Add(1, 2, 3, 4)
	fmt.Println(out)
	out2 := operations.Sub(10, 1, 1, 1)
	fmt.Println(out2)

}
