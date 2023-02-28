package greeter

import "fmt"

// this function is not exported ...
func greeter() {
	fmt.Println("Hello World!")
}

// To export any function the functionName should start with a capital
// letter

func Greeter(name string) string {
	greeter := fmt.Sprintln("Hello", name)
	return greeter
}
